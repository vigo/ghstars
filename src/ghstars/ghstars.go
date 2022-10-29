package ghstars

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/google/go-github/github"
	"github.com/vigo/ghstars/src/version"
	"golang.org/x/oauth2"
)

const defTimeout = 60

// variable declarations.
var (
	client *github.Client
	user   string

	errGitHubAccessTokenRequired = errors.New("please set GITHUB_ACCESS_TOKEN environment variable")
	errInvalidTimeout            = errors.New("invalid timeout value")

	OptVersion      *bool
	OptVerboseMode  *bool
	OptSimpleOutput *bool
	OptTimeout      *int
)

// GHStars represents app structure.
type GHStars struct {
	Out         io.Writer
	AccessToken string
}

func (g *GHStars) flagUsage(code int) func() {
	return func() {
		fmt.Fprintf(
			g.Out,
			usage,
			os.Args[0],
			version.Version,
			defTimeout,
		)
		if code > 0 {
			os.Exit(code)
		}
	}
}

// Run runs the application.
func (g *GHStars) Run() error {
	flag.Usage = g.flagUsage(0)

	OptVersion = flag.Bool(
		"version",
		false,
		fmt.Sprintf("display version information (%s)", version.Version),
	)
	OptVerboseMode = flag.Bool("verbose", false, "verbose/debug mode")

	helpSimpleOutput := "just display the total count and exit"
	OptSimpleOutput = flag.Bool("simple", false, helpSimpleOutput)
	flag.BoolVar(OptSimpleOutput, "s", false, helpSimpleOutput+" (short)")

	helpTimeout := "default timeout in seconds"
	OptTimeout = flag.Int("timeout", defTimeout, helpTimeout)
	flag.IntVar(OptTimeout, "t", defTimeout, helpTimeout+" (short)")

	flag.Parse()

	if *OptVersion {
		fmt.Fprintln(g.Out, version.Version)
		return nil
	}

	if *OptTimeout > 300 || *OptTimeout < 1 {
		return errInvalidTimeout
	}

	timeout := time.Duration(*OptTimeout) * time.Second
	if *OptVerboseMode {
		fmt.Fprintf(g.Out, "[verbose] default timeout is set to: %d seconds\n", *OptTimeout)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if flag.Arg(0) != "" {
		user = flag.Arg(0)
		client = github.NewClient(nil)
		if *OptVerboseMode {
			fmt.Fprintf(g.Out, "[verbose] will fetch w/o authentication, user is %s\n", user)
		}
	} else {
		token, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
		if !ok {
			return errGitHubAccessTokenRequired
		}

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
		if *OptVerboseMode {
			fmt.Fprintln(g.Out, "[verbose] will fetch with using github access token")
		}
	}

	var allRepos []*github.Repository

	if *OptVerboseMode {
		fmt.Fprintln(g.Out, "[verbose] will fetch 100 item per page")
	}

	start := time.Now()

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	if *OptVerboseMode {
		fmt.Fprintln(g.Out, "[verbose] fetching page 1")
	}

	repos, resp, err := client.Repositories.List(ctx, user, opt)
	if err != nil {
		return err // nolint
	}
	allRepos = append(allRepos, repos...)

	if resp.NextPage != 0 && resp.LastPage > 2 {
		if *OptVerboseMode {
			fmt.Fprintf(g.Out, "[verbose] %d page(s) to go, total page count: %d\n", resp.LastPage-1, resp.LastPage)
		}

		var wg sync.WaitGroup
		ch := make(chan []*github.Repository)

		for p := 2; p <= resp.LastPage; p++ {
			wg.Add(1)

			go func(p int) {
				defer wg.Done()

				opts := &github.RepositoryListOptions{
					ListOptions: github.ListOptions{PerPage: 100},
				}
				opts.Page = p

				repos, resp, err := client.Repositories.List(ctx, user, opts)
				if err != nil {
					return
				}
				if *OptVerboseMode {
					fmt.Fprintf(
						g.Out,
						"[verbose] fetching page %d of %d. remaining %d request(s)\n",
						resp.NextPage,
						resp.LastPage,
						resp.Remaining,
					)
				}
				ch <- repos
			}(p)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for repos := range ch {
			allRepos = append(allRepos, repos...)
		}
	}

	elapsed := time.Since(start)

	if *OptVerboseMode {
		fmt.Fprintf(g.Out, "[verbose] fetching took %s seconds\n", elapsed)
	}

	var count int
	result := map[string]int{}

	if user != "" {
		for _, repo := range allRepos {
			if *repo.StargazersCount > 0 && !*repo.Fork {
				result[*repo.FullName] = *repo.StargazersCount
				count += *repo.StargazersCount
			}
		}
	} else {
		for _, repo := range allRepos {
			perms := *repo.Permissions
			if *repo.StargazersCount > 0 && perms["admin"] && !*repo.Fork {
				result[*repo.FullName] = *repo.StargazersCount
				count += *repo.StargazersCount
			}
		}
	}

	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return result[keys[i]] > result[keys[j]]
	})

	// just print total count and exit...
	if *OptSimpleOutput {
		fmt.Fprintln(g.Out, count)
		return nil
	}

	countLen := len(strconv.Itoa(count))

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{},
			{Text: "Name of Repository"},
			{Align: simpletable.AlignCenter, Text: "Star"},
		},
	}

	for i, k := range keys {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%0*d", countLen, i+1)},
			{Text: k},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%0*d", countLen, result[k])},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{},
			{Text: "Number of total star count"},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%0*d", countLen, count)},
		},
	}
	fmt.Fprintln(g.Out, table.String())
	return nil
}

// New instantiates new GHStars instance.
func New(out io.Writer) *GHStars {
	if out == nil {
		out = os.Stdout
	}
	return &GHStars{
		Out: out,
	}
}

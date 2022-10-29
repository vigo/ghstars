package ghstars

var usage = `
usage: [GITHUB_ACCESS_TOKEN=<token>] %[1]s [-flags] [username]

  flags:

  -version           display version information (%s)
  -verbose           verbose/debug mode (default: false)
  -h, -help          display help
  -t, -timeout       default timeout in seconds (default: %d)
  -s, -simple        just display the total count and exit

  examples:

  $ %[1]s <github-username>   star count of <github-username>
  $ %[1]s fatih               star count of fatih
  $ %[1]s tj                  star count of tj

  $ GITHUB_ACCESS_TOKEN="<token>" %[1]s star count of <token> owner

  $ %[1]s -s vigo                 just show count
  $ %[1]s -s fatig                just show count

  $ %[1]s -verbose tj             display in verbose mode
  $ %[1]s -verbose -simple tj     display in verbose mode and just count

  $ %[1]s -t 100 fatih            set timeout to 100 seconds

`

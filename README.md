![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)
![Go](https://img.shields.io/github/go-mod/go-version/vigo/ghstars)
[![Documentation](https://godoc.org/github.com/vigo/ghstars?status.svg)](https://pkg.go.dev/github.com/vigo/ghstars)
[![Go Report Card](https://goreportcard.com/badge/github.com/vigo/ghstars)](https://goreportcard.com/report/github.com/vigo/ghstars)
![Go Build Status](https://github.com/vigo/ghstars/actions/workflows/go.yml/badge.svg)
![GolangCI-Lint Status](https://github.com/vigo/ghstars/actions/workflows/golang-lint.yml/badge.svg)
[![codecov](https://codecov.io/gh/vigo/ghstars/branch/main/graph/badge.svg?token=BTVK8VKVZM)](https://codecov.io/gh/vigo/ghstars)
![Docker Lint Status](https://github.com/vigo/ghstars/actions/workflows/docker-lint.yml/badge.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/vigo/ghstars)
![Docker Size](https://img.shields.io/docker/image-size/vigo/ghstars)
![Docker Build Status](https://github.com/vigo/ghstars/actions/workflows/dockerhub.yml/badge.svg)

# ghstars

Do you want to know how many stars you have? or want to know how many stars
another user has? `ghstars` provides that little functionality for ya!

If you create your own [GitHub Access Token][1]
you can calculate all of your repo stargazers count too!

All you need is to set `GITHUB_ACCESS_TOKEN` environment variable if you want to
fetch all of your repo’s star count!

## Installation

```bash
go install github.com/vigo/ghstars/cmd/ghstars@latest
```

## Usage

You can check command help:

```bash
ghstars -h

ghstars fatih                # public stars count of fatih user
ghstars tj                   # public stars count of fatih tj
ghstars -s vigo              # just how public stars amount of vigo
ghstars -verbose vigo        # public stars count of vigo in debug mode

GITHUB_ACCESS_TOKEN="<your-token>" ghstars  # your repos count
```

How count is done ?

- Public access; star count should be greater than zero and repo fork must be
  false
- Token access; you must be admin in that repo, star count should be greater
  than zero and repo fork must be false

Long story short, **forks are not** counted!

According to [GitHub][2];

- For unauthenticated requests, the rate limit allows for up to **60
  requests** per hour.
- **5000 requests** per hour and per authenticated user.

---

## Docker

https://hub.docker.com/r/vigo/ghstars/

```bash
# latest
docker run vigo/ghstars -h
docker run vigo/ghstars fatih
```

---

## Contributor(s)

* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/ghstars/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

---

## License

This project is licensed under MIT


[1]: https://github.com/settings/tokens/new
[2]: https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting
[coc]: https://github.com/vigo/ghstars/blob/main/CODE_OF_CONDUCT.md

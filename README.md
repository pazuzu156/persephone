# Persephone (Working Title)

[![Build Status](https://img.shields.io/travis/pazuzu156/persephone?style=flat-square)](https://travis-ci.org/pazuzu156/Persephone)
![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability-percentage/pazuzu156/Persephone?label=maintainability&style=flat-square)
![GitHub release](https://img.shields.io/github/release/pazuzu156/persephone?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/pazuzu156/persephone?style=flat-square)
[![Snyk Vulnerabilities for GitHub Repo](https://api.kalebklein.com/persephone/vulns/)](https://app.snyk.io/org/pazuzu156/project/cf386f24-aa5b-4f69-b7ef-657e3f8d3c03)

Persephone is a Discord bot that's used to interface with Lastfm, and is built to be used with the Untrodden Corrodors of Hades Discord server.

## Used Libraries

Libraries used for the bot will be listed here (you can also see them listed in `go.mod`)

- [Disgord](https://github.com/andersfylling/disgord) The Discord library for Go
- [Atlas](https://github.com/pazuzu156/atlas) The command router for Disgord (fork of [Aurora](https://github.com/polaron/aurora))
- [Lastfm-Go](https://github.com/pazuzu156/lastfm-go) The Last.FM API library for Go (fork)
- [gg](https://github.com/fogleman/gg) 2D Image generation for Go
- [genmai](https://github.com/naoina/genmai) Database ORM for Go
- <https://github.com/go-sql-driver/mysql> MySQL driver used with database
- [Colly](https://github.com/gocolly/colly) For hex conversions for embed colors

## Building

Building is quite easy, just make sure you have [Go](https://golang.org/) installed, and run `go build` You should be able to run the compiled executable after making required changes to `config.json` for your bot

### Linux

On Linux, to build for distribution (or to install on your system) use the provided Makefile.

    $ make; make install
    ...

Do not run `make install` with sudo, it will run sudo when needed. Make sure you update your configuration in `~/persephone/config.json` so the bot will run

By default, the makefile will attempt to install the binary in `GOPATH` if that's set, or in `/usr/bin`. To change this, either have `GOPATH` set, or set `PREFIX`

If `GOPATH` is set, but you want to set a custom `PREFIX`, empty `GOPATH`:

    $ GOPATH= PREFIX=/usr/local/bin make install
    ...

Otherwise, just `PREFIX` is fine

## Contributing

Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for more info

## Planned Features

Theres a fair bit I want to do with this bot, I'll add a list of features here when I think of them, and if I find them an obtainable goal

## Donations

Not everyone wants to do a whole lot. But, you can do a small ;) A donation helps keep the bot alive by paying for the server it and it's website run on. And of course, I could always use a beer.

If you want to help out with a small donation, the best ways are the following:

- [PayPal](https://paypal.me/pazuzu156/1)
- Bitcoin: 16vSYHE6Y1icSoFPdc76B32n552YvzZGh6
- Stellar: GCHEI6MJ3QTNEVMK3JB66YT7AHJ7UFTVUY7UDF7TXA3ZGZQOMHVT2AUL
- ZCash: t1eXVKCNwzVYUiT2QS97mr1yBGDbHG2kJfR

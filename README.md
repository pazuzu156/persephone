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
- [Aurora](https://github.com/pazuzu156/aurora) The command router for Disgord (fork)
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

## Contributing

Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for more info

## Planned Features

Theres a fair bit I want to do with this bot, I'll add a list of features here when I think of them, and if I find them an obtainable goal
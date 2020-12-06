# Persephone

![Java w/Gradle](https://github.com/pazuzu156/persephone/workflows/Java%20w/Gradle/badge.svg?branch=java)
![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability-percentage/pazuzu156/Persephone?label=maintainability&style=flat-square)
![GitHub release](https://img.shields.io/github/release/pazuzu156/persephone?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/pazuzu156/persephone?style=flat-square)
[![Snyk Vulnerabilities for GitHub Repo](https://api.kalebklein.com/persephone/vulns/)](https://app.snyk.io/org/pazuzu156/project/cf386f24-aa5b-4f69-b7ef-657e3f8d3c03)

Persephone is a Discord bot that's used to interface with Lastfm, and is built to be used with the Untrodden Corrodors of Hades Discord server.

## Used Libraries

Libraries used for the bot will be listed here (you can also see them listed in `build.gradle`)

- [Gson](https://github.com/google/gson) Google's JSON parsing library
- [JDA](https://github.com/DV8FromTheWorld/JDA) Discord API wrapper for Java
- [JDA-Utilities](https://github.com/JDA-Applications/JDA-Utilities) Extra utilities for JDA (Specifically commands)*
- [Apache Commons Lang3](https://commons.apache.org/proper/commons-lang/) Apache's common language stuff for Java
- [SLF4J](http://www.slf4j.org/) Logging for Java

\* JDA Utilities github releases haven't been updated in years, however the library is still updated. Latest release versions can be found here: https://bintray.com/jagrosh/maven/JDA-Utilities

## Building/Running

Running the bot requires you have a version of Java 14 installed. For Windows, I recommend installing openjdk with chocolatey. Other systems, refer to your distribution's package repositories.

To run, simply edit `src\resources\config.json` with the needed bot configurations, then run using the `run` script for your system.

## Contributing

Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for more info

## Planned Features

As this is a complete bot rewrite in another language, literally everything is a planned feature at this point :rofl:

## Donations

Not everyone wants to do a whole lot. But, you can do a small ;) A donation helps keep the bot alive by paying for the server it and it's website run on. And of course, I could always use a beer.

If you want to help out with a small donation, the best ways are the following:

- [PayPal](https://paypal.me/pazuzu156/1)
- Bitcoin: 16vSYHE6Y1icSoFPdc76B32n552YvzZGh6
- Stellar: GCHEI6MJ3QTNEVMK3JB66YT7AHJ7UFTVUY7UDF7TXA3ZGZQOMHVT2AUL
- ZCash: t1eXVKCNwzVYUiT2QS97mr1yBGDbHG2kJfR

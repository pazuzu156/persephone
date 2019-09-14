## Commands Help {.mdpagehead}

Persephone has a few commands that help retrieve info for things such as current playing tracks, crowns, band info, and more. To use commands, you'll need to follow the bot's commands convention, as laid out below

### Prefix {.mdsechead}

The bot listens to all messages, but will ignore all messages that do not begin with a prefix. This prefix is used to get the bot's attention. The current prefix is a dash: `-`

To use a command, just use the bot's prefix with the command: `-nowplaying`. The bot will then execute the NowPlaying command

### Commands {.mdsechead}

Commands are used to interact with the bot. A command might contain upper case letters within this site's help section, but all commands are case sensitive, and are therefor all built in lowercase. Using upper cases with commands results in your message being ignored.

For help on a given command, you can run `-help [command]` or use the list of bot commands to the left.

### Parameters {.mdsechead}

Some commands have parameters, which will let you do extra stuff with a command. There are two kinds of parameters, `<required>` and `[optional]`. Required parameters are required by the command, and failing to pass it will result in the command either breaking, or simply not running. Optional parameters are optional, and not required to run the command. Optional parameters add some extra stuff to commands if you want to do different things.

For example, if you want to show a specific page for crowns, and for a given user, you could do something like this:

```sh
-crowns @Apollyon#6666 page:2

...
Result of crowns displayed here
...
```

### Aliases {.mdsechead}

Some commands contain aliases. Aliases are basically shorthand commands that execute a command the alias is tied to. This is mostly used for commands that have longer names, so you can type less and do exactly the same as normal.

For example, `-nowplaying` is a pretty long command name, you could use it's alias `-np` and it will run the NowPlaying command.

Aliases can also be used with the Help command. `-help nowplaying` works just the same a `-h np`

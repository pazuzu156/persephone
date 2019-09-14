## BandInfo {.mdpagehead}

The BandInfo command is used to get information on a band. This command grabs info from Last.fm and parses the results in a readable fashion. Since artist images were removed from the Last.fm API, images are retrieved from [Metal Archives](https://www.metal-archives.com/)

### Parameters {.mdsechead}

BandInfo only has one parameter, the band you want to get info on. Passing no parameters will retrieve info on the band you're currently listening to

`[artist]` - Retrieves info on the specified artist

### Aliases {.mdsechead}

BandInfo has three alias: `-bi`, `-ai` and `-artistinfo`

### Examples {.mdsechead}

```sh
-bandinfo darkthrone

...
Darkthrone band info
...
```

```sh
-bandinfo

...
Some artist band info
...
```

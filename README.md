![Logo](assets/logo.png)

[![](https://git.gmantaos.com/haath/Goirate/badges/master/pipeline.svg)](https://git.gmantaos.com/haath/Goirate/pipelines)
[![](https://git.gmantaos.com/haath/Goirate/badges/master/coverage.svg)](https://git.gmantaos.com/haath/Goirate/-/jobs/artifacts/master/browse?job=test)
[![](https://goreportcard.com/badge/git.gmantaos.com/haath/Goirate)](https://goreportcard.com/report/git.gmantaos.com/haath/Goirate)
[![](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![](https://img.shields.io/badge/download-Bintray-blue.svg)](https://dl.bintray.com/gmantaos/Goirate/)

This tool aims to become an all-in-one suite for automating your every pirate-y need.

| <h3>Download</h3> | |
|--------------|-|
| win64 | [ ![Download](https://api.bintray.com/packages/gmantaos/Goirate/win64/images/download.svg) ](https://dl.bintray.com/gmantaos/Goirate/win64) |
| armhf | [ ![Download](https://api.bintray.com/packages/gmantaos/Goirate/armhf/images/download.svg) ](https://dl.bintray.com/gmantaos/Goirate/armhf) |
| aarch64 | [ ![Download](https://api.bintray.com/packages/gmantaos/Goirate/aarch64/images/download.svg) ](https://dl.bintray.com/gmantaos/Goirate/aarch64) |

### 🗺️ P️rogress 

- [x] Pirate Bay scraping for booty
- [x] Command-line plunderin'
- [x] IMDB Scraper
- [x] Robust execution - searching multiple mirrors
- [x] Global configuration management
- [x] Defining series seasons and episodes
- [x] TVDB integration
- [ ] Scanning for new series episodes
- [ ] Defining [sea shanties](https://en.wikipedia.org/wiki/Sea_shanty) and their albums
- [ ] Support for a proxy or VPN to avoid getting flogged
- [ ] Torrent client integration
- [ ] Crontab scanner
    - [ ] Watchlist for single torrents
    - [ ] New series episodes
    - [ ] E-mail notifications
    - [ ] Automatic downloads
    - [ ] RSS Feeds (?)

### Installation

By default [dep](https://github.com/golang/dep) is used for dependency management.

The `Makefile` has a shortcut to running `dep` and `go install`.

```sh
$ make install
```

Using `go get` to fetch dependencies is theoretically possible but it is not
recommended.

## ⚓ Command line tool

### Torrents

The primary source of this tool's torrents is The Pirate Bay.

Commands that search for torrents support the following options.

| | |
|-|-|
| `-j`, `--json` | Output JSON |
| `--mirror "https://pirateproxy.sh/"` | Use a specific pirate bay mirror |
| `--source "https://proxybay.bz/"` | Override default mirror list |
| `--trusted` | Only return torrents whose uploader is either Trusted or VIP |
| `--only-magnet` | Only output magnet links, one on each line |
| `--only-url` | Only output torrent urls, one on each line |
| `-c 7`, `--count 7` | Limit the number of results |

#### Pirate Bay mirrors

To get a list of The Pirate Bay mirrors, use the `goirate mirrors` command.

```shell
$ goirate mirrors
|   | Country |                   URL                    |
|---|---------|------------------------------------------|
| x |   UK    | https://pirateproxy.sh                   |
| x |   NL    | https://thepbproxy.com                   |
| x |   US    | https://thetorrents.red                  |
| x |   US    | https://thetorrents-org.prox.space       |
| x |   US    | https://cruzing.xyz                      |
| x |   US    | https://tpbproxy.nl                      |
| x |   US    | https://thetorrents.rocks                |
| x |   US    | https://proxydl.cf                       |
| x |   US    | https://torrentsblocked.com              |
| x |   US    | https://tpb.crushus.com/thetorrents.org  |
| x |   US    | https://ikwilthetorrents.org             |
| x |   GB    | https://bay.maik.rocks                   |
|   |   FR    | https://www.piratenbaai.ch               |
|   |   US    | https://tpbproxy.gdn                     |
|   |   US    | https://tpb.network                      |
| x |   FR    | https://thetorrents.freeproxy.fun        |
```

By default, the tool will attempt to fetch them from [proxybay.github.io](https://proxybay.github.io). To override that set the `-s` option.

```shell
$ goirate mirrors -s https://proxybay.bz/
```

You can also integrate the tool with any application by getting the output in JSON format using the `--json` flag.

```shell
$ goirate mirrors --json
[
   {
      "url": "https://pirateproxy.sh",
      "country": "uk",
      "status": true
   },
   {
      "url": "https://thepbproxy.com",
      "country": "nl",
      "status": true
   },
   ...
}
```


#### Searching for torrents

```shell
$ goirate search "debian"
                                                  Title                                                      Size    Seeds/Peers
---------------------------------------------------------------------------------------------------------------------------------
 Debian GNU/Linux Bible [ENG] [.pdf]                                                                        7.5 MB   10 / 12
 https://pirateproxy.sh/torrent/** omitted **
---------------------------------------------------------------------------------------------------------------------------------
 Debian 7- System Administration Best Practices, 2013 [PDF]~Storm                                           2.0 MB   9 / 9
 https://pirateproxy.sh/torrent/** omitted **
---------------------------------------------------------------------------------------------------------------------------------
 Debian 9 Stretch minimal install (VirtualBox VDI image)                                                   187.7 MB  6 / 6
 https://pirateproxy.sh/torrent/** omitted **
---------------------------------------------------------------------------------------------------------------------------------
 Debian GNU Linux Bible.zip                                                                                 6.1 MB   2 / 2
 https://pirateproxy.sh/torrent/** omitted **
---------------------------------------------------------------------------------------------------------------------------------
```

## Movies

This tool scrapes [IMDb.com](https://www.imdb.com/) for info on movies.

#### Search

You can fetch a movie - and torrents for it - in three ways.

```sh
# Using its IMDb url
$ goirate movie "https://www.imdb.com/title/tt1028576/"
Secretariat
IMDbID:         1028576
Year:           2010
Rating:         7.1
Duration:       2h 3min
Poster:         https://m.media-amazon.com/images/M/MV5BMTgwNDkyMDU3NV5BMl5BanBnXkFtZTcwNjMyNjI4Mw@@._V1_UX182_CR0,0,182,268_AL_.jpg

Secretariat[2010]DvDrip-aXXo
URL:            ** omitted **
Seeds/Peers:    1 / 1
Size:           735.8 MB
Trusted:        true
Magnet:
magnet:?** omitted **
```

```sh
# Equivalently, using its IMDb ID
$ goirate movie "1028576"
```

```sh
# Using a partial name
$ goirate movie "avengers"
The Avengers
IMDbID:         0848228
Year:           2012
Rating:         8.1
Duration:       2h 23min
...
```

```sh
# Using both a partial name and a release year to narrow down the search
$ goirate movie -y 2018 "avengers"
Avengers: Infinity War
IMDbID:         4154756
Year:           2018
Rating:         8.6
Duration:       2h 29min
...
```

If you don't remember a movie's title or release year very accurately, you can also do a search.

```sh
$ goirate movie-search "harry potter" -c 4
| IMDb ID |             Title              | Year |
|---------|--------------------------------|------|
| 0241527 |      Harry Potter and the      | 2001 |
|         |        Sorcerer's Stone        |      |
| 0330373 | Harry Potter and the Goblet of | 2005 |
|         |              Fire              |      |
| 0417741 |      Harry Potter and the      | 2009 |
|         |       Half-Blood Prince        |      |
| 1201607 |  Harry Potter and the Deathly  | 2011 |
|         |        Hallows: Part 2         |      |
```

## Series

For this tool to manage series, you need to obtain an API key from [TheTVDB.com](https://www.thetvdb.com/)
and include it in Goirate's configuration at `~/.goirate/options.toml`.

Create a watchlist of series, by using the `series add` command.
This stores a list of your series on your account's configuration, along
with the last episode watched for each one. The names can be partial, as they
will be used to search for the full name on the TVDB API. If the last episode is
not specified, the API will be used to fetch the number of the last episode that
aired for this series.

```sh
$ goirate series add "Strike Back" -n "S02E04"
$ goirate series add "The Walking Dead" -n "Season 3 Episode 1"
$ goirate series add "expanse"
```

The `series show` command can be used to display the series currently on the 
watchlist. The `-j` flag also applies here, printing out the JSON format instead.

```sh
$ goirate series show
|   #    |      Series      | Season | Last Episode | Min. Quality |
|--------|------------------|--------|--------------|--------------|
| 280619 | The Expanse      |   3    |      13      |              |
| 153021 | The Walking Dead |   5    |      13      |    1080p     |
```

The `series remove` command can be used to remove a series given either a 
case-insensitive substring in its name, or its TVDB ID.

```sh
$ goirate series remove expanse
$ goirate series remove 153021
```

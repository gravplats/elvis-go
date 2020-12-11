# Elvis

> Elvis is a CLI for generating Spotify playlists

Elvis can generate a Spotify playlist from a few different sources
- From a *setlist* at [setlist.fm](https://www.setlist.fm/)
- From an artist's *top tracks* at [last.fm](https://www.last.fm/)
- From an album list in a [YouTube](https://www.youtube.com/) description (NOTE: highly experimental)

## Build

Elvis is built in [Go](https://golang.org/). Use `make build` to build an executable for the project. If you do not have a local installation of Go, but have [docker](https://www.docker.com/) installed you can use `make docker-build` which will build an executable for `darwin/amd64` (macOS). The executable will be located in the `dist` folder.

## Commands

```
Elvis is a CLI for generating Spotify playlists from various web APIs.

Usage:
  elvis [command]

Available Commands:
  help        Help about any command
  init        Generate a config file in the home directory (will not overwrite)
  setlist     Generate a playlist from a setlist.fm setlist
  top         Generate a playlist from the last.fm top tracks for an artist
  yt          Generate a playlist from albums in a YouTube description

Flags:
  -h, --help   help for elvis

Use "elvis [command] --help" for more information about a command.
```

### `init`

```
Generate a config file in the home directory (will not overwrite)

Usage:
  elvis init [flags]

Examples:
  elvis init

Flags:
  -h, --help   help for init
```

Run `elvis init` to create a configuration file (`.elvis.json`) in your home folder.

```json
{
  "lastfm": {
    "credentials": {
      "key": "",
      "secret": ""
    }
  },
  "setlistfm": {
    "credentials": {
      "key": ""
    }
  },
  "spotify": {
    "credentials": {
      "id": "",
      "secret": ""
    }
  },
  "youtube": {
    "credentials": {
      "key": ""
    }
  }
}
```

All `elvis` commands use the Spotify Web API to generate Spotify playlists. In order to access the API you need to generate a *client id* and *client secret* over at the Spotify for Developers [dashboard](https://developer.spotify.com/dashboard/). Add these keys to `spotify.credentials.id` and `spotify.credentials.secret` respectively in the configuration file.

### `setlist`

```
Generate a playlist from a setlist.fm setlist

Usage:
  elvis setlist <value> [flags]

Examples:
  elvis setlist https://www.setlist.fm/setlist/opeth/2017/finlandia-talo-helsinki-finland-53e3dba9.html
  elvis setlist 53e3dba9

Flags:
  -h, --help   help for setlist
```

Use `elvis setlist <value>` to generate a playlist from a setlist at setlist.fm. `<value>` can be either a setlist ID or a URL to a setlist. The `setlist` command uses the [setlist.fm API](https://api.setlist.fm/docs/1.0/index.html). In order to use the setlist.fm API you need to sign up for an API key. Add this key to `setlistfm.credentials.key` in the configuration file.

### `top`

```
Generate a playlist from the last.fm top tracks for an artist

Usage:
  elvis top <artist> [flags]

Examples:
  elvis top opeth

Flags:
  -h, --help        help for top
      --limit int   Number of tracks (default 10)
```

Use `elvis top <artist>` to generate a playlist from the artist's top tracks at last.fm. Use `--limit <value>` to specify the number of tracks to include. The default value is 10. The `top` command uses the [last.fm API](https://www.last.fm/api). In order to use the last.fm API you need to sign up to get a *client key* and *client secret*. Add these keys to `lastfm.credentials.key` and `lastfm.credentials.secret` respectively in the configuration file.

### `yt`

```
Generate a playlist from albums in a YouTube description

Usage:
  elvis yt <value> [flags]

Examples:
  elvis yt https://www.youtube.com/watch?v=aEdho6H0hFY
  elvis yt aEdho6H0hFY

Flags:
  -h, --help   help for yt
```

Use `elvis yt <value>` to generate a playlist from an album list in a YouTube description (NOTE: highly experimental). `<value>` can either be a YouTube video ID or a URL to a YouTube video. The `yt` command uses the [YouTube API](https://developers.google.com/youtube/v3). In order to use the YouTube API you need to sign up for an API key. Add this key to `youtube.credentials.key` in the configuration file.
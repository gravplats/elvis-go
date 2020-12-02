package main

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/cmd"
	"github.com/mrydengren/elvis/pkg/config"
	"log"
	"os"
	"strings"
)

const (
	SPOTIFY_ID     = "SPOTIFY_ID"
	SPOTIFY_SECRET = "SPOTIFY_SECRET"
)

const Usage = `Elvis is a CLI for generating Spotify playlists from various web APIs.

Usage:
   elvis setlist <value>
   elvis top <artist>

Arguments:
   artist   the name of the artist
   value    either the setlist ID or the setlist URL

Examples:
   elvis setlist https://www.setlist.fm/setlist/opeth/2017/finlandia-talo-helsinki-finland-53e3dba9.html
   elvis setlist 53e3dba9

   elvis top opeth`

func main() {
	cfg, err := config.Read(".elvis.json")
	if err != nil {
		log.Fatal(err)
	}

	// The lastfm API wrapper reads these, if present.
	os.Setenv(lastfm.ApiKey, cfg.Lastfm.Key)
	os.Setenv(lastfm.ApiSecret, cfg.Lastfm.Secret)
	// The setlistfm API wrapper reads these, if present.
	os.Setenv(setlistfm.ApiKey, cfg.Setlistfm.Key)
	// The Spotify API wrapper reads these, if present.
	os.Setenv(SPOTIFY_ID, cfg.Spotify.Id)
	os.Setenv(SPOTIFY_SECRET, cfg.Spotify.Secret)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(Usage)
		os.Exit(0)
	}

	switch args[0] {
	case "setlist":
		if len(args) < 2 {
			fmt.Println("Missing <value>. See help for more information.")
			os.Exit(0)
		}

		value := args[1]
		cmd.Setlist(value)
	case "top":
		if len(args) < 2 {
			fmt.Println("Missing <artist>. See help for more information.")
			os.Exit(0)
		}

		artist := strings.Join(args[1:], " ")
		// TODO: limit should be a CLI flag
		limit := 10

		cmd.Top(artist, limit)
	case "help":
		fallthrough
	default:
		fmt.Println(Usage)
		os.Exit(0)
	}
}

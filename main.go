package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/cmd"
	"github.com/mrydengren/elvis/pkg/config"
	"github.com/spf13/viper"
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
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName(".elvis")
	viper.SetConfigType("json")

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to read config file.")
	}

	var cfg config.Config
	viper.Unmarshal(&cfg)

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

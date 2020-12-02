package main

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/cmd"
	"github.com/mrydengren/elvis/pkg/config"
	"log"
	"os"
)

const (
	SPOTIFY_ID     = "SPOTIFY_ID"
	SPOTIFY_SECRET = "SPOTIFY_SECRET"
)

func main() {
	cfg, err := config.Read(".elvis.json")
	if err != nil {
		log.Fatal(err)
	}

	// The lastfm API wrapper reads these, if present.
	os.Setenv(lastfm.ApiKey, cfg.Lastfm.Key)
	os.Setenv(lastfm.ApiSecret, cfg.Lastfm.Secret)
	// The Spotify API wrapper reads these, if present.
	os.Setenv(SPOTIFY_ID, cfg.Spotify.Id)
	os.Setenv(SPOTIFY_SECRET, cfg.Spotify.Secret)

	cmd.Top("Opeth", 10)
}

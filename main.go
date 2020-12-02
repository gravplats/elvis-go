package main

import (
	"github.com/mrydengren/elvis/pkg/config"
	"github.com/mrydengren/elvis/pkg/playlist"
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

	// The Spotify API wrapper reads these, if present.
	os.Setenv(SPOTIFY_ID, cfg.Spotify.Id)
	os.Setenv(SPOTIFY_SECRET, cfg.Spotify.Secret)

	tracks := []playlist.Track{
		playlist.Track{
			Artist: "Opeth",
			Name:   "Heir Apparent",
		},
		playlist.Track{
			Artist: "Opeth",
			Name:   "Deliverance",
		},
		playlist.Track{
			Artist: "Opeth",
			Name:   "Ghost of Perdition",
		},
	}

	tracklist := playlist.Tracklist{
		Artist: "Opeth",
		Id:     "opeth",
		Tracks: tracks,
	}

	playlist.Create(tracklist)
}

package main

import (
	"github.com/mrydengren/elvis/pkg/playlist"
	"log"
	"os"
)

const (
	SPOTIFY_ID     = "SPOTIFY_ID"
	SPOTIFY_SECRET = "SPOTIFY_SECRET"
)

func main() {
	if os.Getenv(SPOTIFY_ID) == "" {
		log.Fatalf("Missing env %s", SPOTIFY_ID)
	}
	if os.Getenv(SPOTIFY_SECRET) == "" {
		log.Fatalf("Missing env %s", SPOTIFY_SECRET)
	}

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

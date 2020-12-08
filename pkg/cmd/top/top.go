package top

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
)

func Top(artist string, limit int) error {
	spinner.Start("Fetching top tracks.")

	client := lastfm.NewClient()
	options := &lastfm.TopTracksOptions{
		Limit: &limit,
	}

	toptracks, err := client.ArtistTopTracks(artist, options)
	if err != nil {
		spinner.Fail()
		return err
	}

	debug.DumpJson(toptracks, "lastfm-toptracks.json")

	itemGroup := fromTopTracks(toptracks)

	spinner.Succeed()

	return playlist.Create(itemGroup)
}

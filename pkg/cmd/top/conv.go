package top

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/playlist"
)

func FromTopTracks(toptracks *lastfm.TopTracksWrapper) playlist.SearchItemGroup {
	var items []playlist.SearchItem

	for _, tt := range toptracks.Toptracks.Track {
		item := playlist.SearchItem{
			Artist: tt.Artist.Name,
			Name:   tt.Name,
		}

		items = append(items, item)
	}

	return playlist.SearchItemGroup{
		Artist: toptracks.Toptracks.Attr.Artist,
		Items:  items,
		Type:   playlist.SearchItemTypeTrack,
	}
}

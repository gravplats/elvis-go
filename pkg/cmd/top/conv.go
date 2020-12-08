package top

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/playlist"
)

func fromTopTracks(toptracks *lastfm.TopTracksWrapper) playlist.ItemGroup {
	var items []playlist.Item

	for _, tt := range toptracks.Toptracks.Track {
		item := playlist.Item{
			Artist: tt.Artist.Name,
			Name:   tt.Name,
		}

		items = append(items, item)
	}

	return playlist.ItemGroup{
		Name:  toptracks.Toptracks.Attr.Artist,
		Items: items,
		Type:  playlist.ItemGroupTypeTrack,
	}
}

package setlist

import (
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/playlist"
)

func FromSetlist(setlist *setlistfm.Setlist) playlist.SearchItemGroup {
	var items []playlist.SearchItem

	for _, set := range setlist.Sets.Set {
		for _, song := range set.Song {
			item := playlist.SearchItem{
				Name: song.Name,
			}

			if song.Cover.Name != "" {
				item.Artist = song.Cover.Name
			} else {
				item.Artist = setlist.Artist.Name
			}

			items = append(items, item)
		}
	}

	return playlist.SearchItemGroup{
		Artist: setlist.Artist.Name,
		Items:  items,
		Type:   playlist.SearchItemTypeTrack,
	}
}

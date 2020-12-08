package setlist

import (
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/playlist"
)

func fromSetlist(setlist *setlistfm.Setlist) playlist.ItemGroup {
	var items []playlist.Item

	for _, set := range setlist.Sets.Set {
		for _, song := range set.Song {
			item := playlist.Item{
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

	return playlist.ItemGroup{
		Name:  setlist.Artist.Name,
		Items: items,
		Type:  playlist.ItemGroupTypeTrack,
	}
}

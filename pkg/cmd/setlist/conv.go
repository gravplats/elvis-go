package setlist

import (
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/playlist"
)

func FromSetlist(setlist *setlistfm.Setlist) playlist.Tracklist {
	var tracks []playlist.Track

	for _, set := range setlist.Sets.Set {
		for _, song := range set.Song {
			track := playlist.Track{
				Name: song.Name,
			}

			if song.Cover.Name != "" {
				track.Artist = song.Cover.Name
			} else {
				track.Artist = setlist.Artist.Name
			}

			tracks = append(tracks, track)
		}
	}

	return playlist.Tracklist{
		Artist: setlist.Artist.Name,
		Id:     setlist.Id,
		Tracks: tracks,
	}
}

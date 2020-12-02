package playlist

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"strings"
)

func FromSetlist(setlist *setlistfm.Setlist) Tracklist {
	var tracks []Track

	for _, set := range setlist.Sets.Set {
		for _, song := range set.Song {
			track := Track{
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

	return Tracklist{
		Artist: setlist.Artist.Name,
		Id:     setlist.Id,
		Tracks: tracks,
	}
}

func FromTopTracks(toptracks *lastfm.TopTracksWrapper) Tracklist {
	var tracks []Track

	for _, tt := range toptracks.Toptracks.Track {
		track := Track{
			Artist: tt.Artist.Name,
			Name:   tt.Name,
		}

		tracks = append(tracks, track)
	}

	return Tracklist{
		Artist: toptracks.Toptracks.Attr.Artist,
		Id:     strings.ToLower(toptracks.Toptracks.Attr.Artist),
		Tracks: tracks,
	}
}

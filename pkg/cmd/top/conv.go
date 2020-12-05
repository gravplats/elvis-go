package top

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/playlist"
	"strings"
)

func FromTopTracks(toptracks *lastfm.TopTracksWrapper) playlist.Tracklist {
	var tracks []playlist.Track

	for _, tt := range toptracks.Toptracks.Track {
		track := playlist.Track{
			Artist: tt.Artist.Name,
			Name:   tt.Name,
		}

		tracks = append(tracks, track)
	}

	return playlist.Tracklist{
		Artist: toptracks.Toptracks.Attr.Artist,
		Id:     strings.ToLower(toptracks.Toptracks.Attr.Artist),
		Tracks: tracks,
	}
}

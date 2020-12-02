package playlist

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"strings"
)

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

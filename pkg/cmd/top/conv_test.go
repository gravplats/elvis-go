package top

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/playlist"
	"reflect"
	"testing"
)

func TestFromTopTracks(t *testing.T) {
	toptracks := lastfm.TopTracksWrapper{
		Toptracks: lastfm.TopTracks{
			Track: []lastfm.TopTrack{
				lastfm.TopTrack{
					Artist: lastfm.TopTrackArtist{
						Name: "Opeth",
					},
					Name: "Heir Apparent",
				},
			},
		},
	}

	got := FromTopTracks(&toptracks)

	want := playlist.Tracklist{
		Artist: "Opeth",
		Id:     "opeth",
		Tracks: []playlist.Track{
			playlist.Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
		},
	}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

package playlist

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
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

	want := Tracklist{
		Artist: "Opeth",
		Id:     "opeth",
		Tracks: []Track{
			Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
		},
	}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

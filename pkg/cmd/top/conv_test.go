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

	want := playlist.SearchItemGroup{
		Artist: "Opeth",
		Items: []playlist.SearchItem{
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
		},
		Type: playlist.SearchItemTypeTrack,
	}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

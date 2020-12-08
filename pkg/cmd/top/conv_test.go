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

	got := fromTopTracks(&toptracks)

	want := playlist.ItemGroup{
		Name: "Opeth",
		Items: []playlist.Item{
			playlist.Item{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
		},
		Type: playlist.ItemGroupTypeTrack,
	}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

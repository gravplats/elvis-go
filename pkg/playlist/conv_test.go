package playlist

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"reflect"
	"testing"
)

func TestFromSetlistWithMultipleSongs(t *testing.T) {
	artist := setlistfm.SetlistArtist{
		Name: "Opeth",
	}

	set := []setlistfm.SetlistSet{
		setlistfm.SetlistSet{
			Song: []setlistfm.SetlistSong{
				setlistfm.SetlistSong{
					Name: "Heir Apparent",
				},
				setlistfm.SetlistSong{
					Name: "Ghost of Perdition",
				},
			},
		},
	}

	sets := setlistfm.SetlistSets{
		Set: set,
	}

	setlist := setlistfm.Setlist{
		Artist: artist,
		Id:     "123abc",
		Sets:   sets,
	}

	want := Tracklist{
		Artist: "Opeth",
		Id:     "123abc",
		Tracks: []Track{
			Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			Track{
				Artist: "Opeth",
				Name:   "Ghost of Perdition",
			},
		},
	}

	got := FromSetlist(&setlist)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestFromSetlistWithCover(t *testing.T) {
	artist := setlistfm.SetlistArtist{
		Name: "Opeth",
	}

	set := []setlistfm.SetlistSet{
		setlistfm.SetlistSet{
			Song: []setlistfm.SetlistSong{
				setlistfm.SetlistSong{
					Name: "Heir Apparent",
				},
				setlistfm.SetlistSong{
					Cover: setlistfm.SetlistArtist{
						Name: "Napalm Death",
					},
					Name: "You Suffer",
				},
			},
		},
	}

	sets := setlistfm.SetlistSets{
		Set: set,
	}

	setlist := setlistfm.Setlist{
		Artist: artist,
		Id:     "123abc",
		Sets:   sets,
	}

	want := Tracklist{
		Artist: "Opeth",
		Id:     "123abc",
		Tracks: []Track{
			Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			Track{
				Artist: "Napalm Death",
				Name:   "You Suffer",
			},
		},
	}

	got := FromSetlist(&setlist)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestFromSetlistWithMultipleSets(t *testing.T) {
	artist := setlistfm.SetlistArtist{
		Name: "Opeth",
	}

	set := []setlistfm.SetlistSet{
		setlistfm.SetlistSet{
			Song: []setlistfm.SetlistSong{
				setlistfm.SetlistSong{
					Name: "Heir Apparent",
				},
			},
		},
		setlistfm.SetlistSet{
			Song: []setlistfm.SetlistSong{
				setlistfm.SetlistSong{
					Name: "Deliverance",
				},
			},
		},
	}

	sets := setlistfm.SetlistSets{
		Set: set,
	}

	setlist := setlistfm.Setlist{
		Artist: artist,
		Id:     "123abc",
		Sets:   sets,
	}

	want := Tracklist{
		Artist: "Opeth",
		Id:     "123abc",
		Tracks: []Track{
			Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			Track{
				Artist: "Opeth",
				Name:   "Deliverance",
			},
		},
	}

	got := FromSetlist(&setlist)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

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

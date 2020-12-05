package setlist

import (
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/playlist"
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

	want := playlist.Tracklist{
		Artist: "Opeth",
		Id:     "123abc",
		Tracks: []playlist.Track{
			playlist.Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			playlist.Track{
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

	want := playlist.Tracklist{
		Artist: "Opeth",
		Id:     "123abc",
		Tracks: []playlist.Track{
			playlist.Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			playlist.Track{
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

	want := playlist.Tracklist{
		Artist: "Opeth",
		Id:     "123abc",
		Tracks: []playlist.Track{
			playlist.Track{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			playlist.Track{
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

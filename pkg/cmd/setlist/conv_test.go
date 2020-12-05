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
		Sets:   sets,
	}

	want := playlist.SearchItemGroup{
		Artist: "Opeth",
		Items: []playlist.SearchItem{
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Ghost of Perdition",
			},
		},
		Type: playlist.SearchItemTypeTrack,
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
		Sets:   sets,
	}

	want := playlist.SearchItemGroup{
		Artist: "Opeth",
		Items: []playlist.SearchItem{
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			playlist.SearchItem{
				Artist: "Napalm Death",
				Name:   "You Suffer",
			},
		},
		Type: playlist.SearchItemTypeTrack,
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
		Sets:   sets,
	}

	want := playlist.SearchItemGroup{
		Artist: "Opeth",
		Items: []playlist.SearchItem{
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Deliverance",
			},
		},
		Type: playlist.SearchItemTypeTrack,
	}

	got := FromSetlist(&setlist)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

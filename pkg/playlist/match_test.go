package playlist

import (
	"reflect"
	"testing"
)

func TestGetBestMatch(t *testing.T) {
	item := SearchItem{
		Artist: "Opeth",
		Name:   "Heir Apparent",
	}

	resources := []Resource{
		Resource{
			ID:   "1",
			Name: "Heir Apparent",
		},
	}

	want := Resource{
		ID:   "1",
		Name: "Heir Apparent",
	}

	got := getBestMatch(item, resources)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestMatch(t *testing.T) {
	group := SearchItemGroup{
		Artist: "Opeth",
		Items: []SearchItem{
			SearchItem{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			SearchItem{
				Artist: "Opeth",
				Name:   "Ghost of Perdition",
			},
		},
	}

	resources := [][]Resource{
		[]Resource{
			Resource{
				ID:   "1",
				Name: "Heir Apparent",
			},
		},
		[]Resource{
			Resource{
				ID:   "2",
				Name: "Ghost of Perdition",
			},
		},
	}

	want := []SearchMatch{
		SearchMatch{
			Artist: "Opeth",
			ID:     "1",
			Name:   "Heir Apparent",
		},
		SearchMatch{
			Artist: "Opeth",
			ID:     "2",
			Name:   "Ghost of Perdition",
		},
	}

	got := match(group, resources)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

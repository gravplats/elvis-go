package playlist

import (
	"reflect"
	"testing"
)

func TestGetBestMatch(t *testing.T) {
	item := Item{
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
	group := ItemGroup{
		Name: "Opeth",
		Items: []Item{
			Item{
				Artist: "Opeth",
				Name:   "Heir Apparent",
			},
			Item{
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

	want := []Match{
		Match{
			Artist: "Opeth",
			ID:     "1",
			Name:   "Heir Apparent",
		},
		Match{
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

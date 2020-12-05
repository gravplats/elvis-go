package playlist

import (
	"github.com/zmb3/spotify"
	"reflect"
	"testing"
)

func TestSliceMaxMinus1(t *testing.T) {
	values := []spotify.ID{"1"}

	got := Split(values, 2)

	want := [][]spotify.ID{
		[]spotify.ID{"1"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestSliceMax(t *testing.T) {
	values := []spotify.ID{"1", "2"}

	got := Split(values, 2)

	want := [][]spotify.ID{
		[]spotify.ID{"1", "2"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestSliceMaxPlus1(t *testing.T) {
	values := []spotify.ID{"1", "2", "3"}

	got := Split(values, 2)

	want := [][]spotify.ID{
		[]spotify.ID{"1", "2"},
		[]spotify.ID{"3"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

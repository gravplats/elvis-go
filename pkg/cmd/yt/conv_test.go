package yt

import (
	"github.com/mrydengren/elvis/pkg/playlist"
	"reflect"
	"testing"
)

func TestFromDescription(t *testing.T) {
	type fixture struct {
		name  string
		input string
	}

	fixtures := []fixture{
		fixture{
			name: "first line ending in colon - \"artist- album title\" on same line",
			input: "" +
				"Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans:\n" +
				"Opeth- Still Life\n" +
				"https://opeth.com/albums/still-life",
		},
		fixture{
			name: "first line ending in colon plus whitespace - \"artist- album title\" on same line",
			input: "" +
				"Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans:   \n" +
				"Opeth- Still Life\n" +
				"https://opeth.com/albums/still-life",
		},
		fixture{
			name: "first line ending in colon - artist album title on separate lines",
			input: "" +
				"Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans:\n" +
				"Opeth\n" +
				"Still Life\n" +
				"https://opeth.com/albums/still-life",
		},
		fixture{
			name: "first line ending in colon plus whitespace - artist album title on separate lines",
			input: "" +
				"Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans:   \n" +
				"Opeth\n" +
				"Still Life\n" +
				"https://opeth.com/albums/still-life",
		},
		fixture{
			name: "empty line before albums",
			input: "" +
				"Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans\n" +
				"\n" +
				"Opeth\n" +
				"Still Life\n" +
				"https://opeth.com/albums/still-life",
		},
	}

	want := playlist.SearchItemGroup{
		Artist: "Opeth",
		Items: []playlist.SearchItem{
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Still Life",
			},
		},
		Type: playlist.SearchItemTypeAlbum,
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got := FromDescription("Opeth", fixture.input)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %+v, want: %+v", got, want)
			}
		})
	}
}

func TestFromDescriptionMultipleAlbums(t *testing.T) {
	description := "" +
		"Opeth\n" +
		"Still Life\n" +
		"https://opeth.com/albums/still-life" +
		"\n" +
		"Opeth\n" +
		"Blackwater Park\n" +
		"https://opeth.com/albums/blackwater-park"

	want := playlist.SearchItemGroup{
		Artist: "Opeth",
		Items: []playlist.SearchItem{
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Still Life",
			},
			playlist.SearchItem{
				Artist: "Opeth",
				Name:   "Blackwater Park",
			},
		},
		Type: playlist.SearchItemTypeAlbum,
	}

	got := FromDescription("Opeth", description)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestEndsWithColon(t *testing.T) {
	type fixture struct {
		name  string
		input string
		want  bool
	}

	fixtures := []fixture{
		fixture{
			name:  "colon",
			input: "Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans:",
			want:  true,
		},
		fixture{
			name:  "colon + whitespace",
			input: "Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans:   ",
			want:  true,
		},
		fixture{
			name:  "text",
			input: "Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans",
			want:  false,
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got := endsInColon(fixture.input)
			if got != fixture.want {
				t.Errorf("got: %+v, want: %+v", got, fixture.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type fixture struct {
		name  string
		input string
		want  bool
	}

	fixtures := []fixture{
		fixture{
			name:  "empty",
			input: "",
			want:  true,
		},
		fixture{
			name:  "whitespace",
			input: "   ",
			want:  true,
		},
		fixture{
			name:  "text",
			input: "Cupcake ipsum dolor sit amet chocolate cake soufflé chocolate jelly beans",
			want:  false,
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got := isEmpty(fixture.input)
			if got != fixture.want {
				t.Errorf("got: %+v, want: %+v", got, fixture.want)
			}
		})
	}
}

func TestStartsWithHttpProtocol(t *testing.T) {
	type fixture struct {
		name  string
		input string
		want  bool
	}

	fixtures := []fixture{
		fixture{
			name:  "http",
			input: "http://",
			want:  true,
		},
		fixture{
			name:  "https",
			input: "https://",
			want:  true,
		},
		fixture{
			name:  "text",
			input: "Cupcake ipsum dolor sit amet http://",
			want:  false,
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got := startsWithHttpProtocol(fixture.input)
			if got != fixture.want {
				t.Errorf("got: %+v, want: %+v", got, fixture.want)
			}
		})
	}
}

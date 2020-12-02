package config

import (
	"reflect"
	"testing"
)

func TestReadValidConfigFile(t *testing.T) {
	want := Config{
		Lastfm: LastfmApiKeys{
			Key:    "key value",
			Secret: "secret value",
		},
		Spotify: SpotifyApiKeys{
			Id:     "id value",
			Secret: "secret value",
		},
	}

	got, err := Read("fixtures/config.json")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestReadNonJsonFile(t *testing.T) {
	want := ErrNotValidJson
	_, got := Read("fixtures/config-empty")

	if got != want {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func TestReadEmptyJson(t *testing.T) {
	_, err := Read("fixtures/config-empty.json")
	if err == nil {
		t.Errorf("Expected error but got %+v", err)
	}

	e := Error{
		Errors: []error{
			ErrMissingLastfmKey,
			ErrMissingLastfmSecret,
			ErrMissingSpotifyId,
			ErrMissingSpotifySecret,
		},
	}

	got := err.Error()
	want := e.Error()

	if got != want {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type LastfmApiKeys struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type SetlistfmApiKeys struct {
	Key string `json:"key"`
}

type SpotifyApiKeys struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}

type Config struct {
	Lastfm    LastfmApiKeys    `json:"lastfm"`
	Setlistfm SetlistfmApiKeys `json:"setlistfm"`
	Spotify   SpotifyApiKeys   `json:"spotify"`
}

type Error struct {
	Errors []error
}

var (
	ErrNotValidJson         = errors.New("not valid JSON")
	ErrMissingLastfmKey     = errors.New("missing lastfm key")
	ErrMissingLastfmSecret  = errors.New("missing lastfm secret")
	ErrMissingSetlistfmKey  = errors.New("missing setlistfm key")
	ErrMissingSpotifyId     = errors.New("missing spotify id")
	ErrMissingSpotifySecret = errors.New("missing spotify secret")
)

func (e Error) Error() string {
	var elems = []string{"One or more config settings are missing:"}
	for _, v := range e.Errors {
		elems = append(elems, fmt.Sprintf("\t- %s", v.Error()))
	}
	return strings.Join(elems, "\n")
}

func Read(dir string) (Config, error) {
	f, err := os.Open(dir)
	if err != nil {
		return Config{}, err
	}

	d := json.NewDecoder(f)

	var cfg Config
	err = d.Decode(&cfg)
	if err != nil {
		if err == io.EOF {
			return Config{}, ErrNotValidJson
		}
		return Config{}, err
	}

	// The config could be validated using JSON schema validation instead,
	// but let us stick to the standard packages for now.
	var errs []error

	if cfg.Lastfm.Key == "" {
		errs = append(errs, ErrMissingLastfmKey)
	}
	if cfg.Lastfm.Secret == "" {
		errs = append(errs, ErrMissingLastfmSecret)
	}
	if cfg.Setlistfm.Key == "" {
		errs = append(errs, ErrMissingSetlistfmKey)
	}
	if cfg.Spotify.Id == "" {
		errs = append(errs, ErrMissingSpotifyId)
	}
	if cfg.Spotify.Secret == "" {
		errs = append(errs, ErrMissingSpotifySecret)
	}
	if len(errs) > 0 {
		return Config{}, Error{Errors: errs}
	}
	return cfg, nil
}

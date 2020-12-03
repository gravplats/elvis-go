package config

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

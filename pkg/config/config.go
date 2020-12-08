package config

type LastfmApiKeys struct {
	Credentials struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"credentials"`
}

type SetlistfmApiKeys struct {
	Credentials struct {
		Key string `json:"key"`
	} `json:"credentials"`
}

type SpotifyApiKeys struct {
	Credentials struct {
		Id     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"credentials"`
}

type YouTubeApiKeys struct {
	Credentials struct {
		Key string `json:"key""`
	} `json:"credentials"`
}

type Config struct {
	Lastfm    LastfmApiKeys    `json:"lastfm"`
	Setlistfm SetlistfmApiKeys `json:"setlistfm"`
	Spotify   SpotifyApiKeys   `json:"spotify"`
	YouTube   YouTubeApiKeys   `json:"youtube"`
}

package lastfm

import (
	"log"
	"net/url"
	"strconv"
)

type TopTrackArtist struct {
	Mbid string `json:"mbid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type TopTrackAttr struct {
	Rank string `json:"string"`
}

type TopTrackImage struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

type TopTrack struct {
	Artist     TopTrackArtist  `json:"artist"`
	Attr       TopTrackAttr    `json:"@attr"`
	Image      []TopTrackImage `json:"image"`
	Listeners  string          `json:"listeners"`
	Mbid       string          `json:"mbid"`
	Name       string          `json:"name"`
	Playcount  string          `json:"playcount"`
	Streamable string          `json:"streamable"`
	Url        string          `json:"url"`
}

type TopTracksAttr struct {
	Artist     string `json:"artist"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	Total      string `json:"total"`
	TotalPages string `json:"totalPages"`
}

type TopTracks struct {
	Attr  TopTracksAttr `json:"@attr"`
	Track []TopTrack    `json:"track"`
}

type TopTracksWrapper struct {
	Toptracks TopTracks `json:"toptracks"`
}

type TopTracksOptions struct {
	Autocorrect *bool
	Limit       *int
	Mbid        *string
	Page        *int
}

func (c *Client) ArtistTopTracks(artist string, opt *TopTracksOptions) (*TopTracksWrapper, error) {
	values := url.Values{}
	values.Set("artist", artist)

	if opt != nil {
		if opt.Autocorrect != nil {
			values.Set("autocorrect", parseBool(*opt.Autocorrect))
		}
		if opt.Limit != nil {
			values.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Mbid != nil {
			values.Set("mbid", *opt.Mbid)
		}
		if opt.Page != nil {
			values.Set("page", strconv.Itoa(*opt.Page))
		}
	}

	var toptracks TopTracksWrapper

	err := c.get(MethodArtistTopTracks, values, &toptracks)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &toptracks, nil
}

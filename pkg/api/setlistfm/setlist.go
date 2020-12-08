package setlistfm

type SetlistArtist struct {
	Disambiguation string `json:"disambiguation"`
	Mbid           string `json:"mbid"`
	Name           string `json:"name"`
	SortName       string `json:"sortName"`
	Tmid           string `json:"tmid"`
	Url            string `json:"url"`
}

type SetlistSet struct {
	Encore int64         `json:"encore"`
	Name   string        `json:"name"`
	Song   []SetlistSong `json:"song"`
}

type SetlistSets struct {
	Set []SetlistSet `json:"set"`
}

type SetlistSong struct {
	Cover SetlistArtist `json:"cover"`
	Info  string        `json:"info"`
	Name  string        `json:"name"`
	Tape  bool          `json:"tape"`
}

type SetlistTour struct {
	Name string `json:"name"`
}

type SetlistVenueCity struct {
	Coords struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"coords"`
	Country struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:""`
	Id        string `json:"id"`
	Name      string `json:"name"`
	State     string `json:"state"`
	StateCode string `json:"stateCode"`
}

type SetlistVenue struct {
	City SetlistVenueCity `json:"city"`
	Id   string           `json:"id"`
	Name string           `json:"name"`
	Url  string           `json:"url"`
}

type Setlist struct {
	Artist      SetlistArtist `json:"artist"`
	EventDate   string        `json:"eventDate"`
	Id          string        `json:"id"`
	LastUpdated string        `json:"lastUpdated"`
	Sets        SetlistSets   `json:"sets"`
	Tour        SetlistTour   `json:"tour"`
	Url         string        `json:"url"`
	Venue       SetlistVenue  `json:"venue"`
	VersionId   string        `json:"versionId"`
}

func (c *Client) Setlist(setlistId string) (*Setlist, error) {
	u := c.baseUrl + "setlist/" + setlistId

	var setlist Setlist
	err := c.get(u, &setlist)
	if err != nil {
		return nil, err
	}

	return &setlist, nil
}

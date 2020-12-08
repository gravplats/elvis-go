package lastfm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const baseAddress = "http://ws.audioscrobbler.com/2.0/"

const (
	ApiKey    = "LASTFM_KEY"
	ApiSecret = "LASTFM_SECRET"
)

const (
	MethodArtistTopTracks = "artist.gettoptracks"
)

type Client struct {
	accept    string
	apiKey    string
	apiSecret string
	baseUrl   string
	http      *http.Client
}

func NewClient() Client {
	return Client{
		accept:    "json",
		apiKey:    os.Getenv(ApiKey),
		apiSecret: os.Getenv(ApiSecret),
		baseUrl:   baseAddress,
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) AcceptJson() {
	c.accept = "json"
}

func (c *Client) AcceptXml() {
	c.accept = "xml"
}

func (c *Client) SetAuthInfo(apiKey string, apiSecret string) {
	c.apiKey = apiKey
	c.apiSecret = apiSecret
}

func (c *Client) get(method string, values url.Values, result interface{}) error {
	values.Add("api_key", c.apiKey)
	values.Add("format", c.accept)
	values.Add("method", method)

	u := baseAddress + "?" + values.Encode()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// TODO: decode error in response.
		return fmt.Errorf("HTTP NOT OK")
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

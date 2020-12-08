package lastfm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"error"`
}

func (e Error) Error() string {
	return e.Message
}

func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("lastfm: HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("lastfm: couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	return fmt.Errorf("lastfm: %s", e)
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
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

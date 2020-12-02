package setlistfm

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

const baseAddress = "https://api.setlist.fm/rest/1.0/"

const (
	ApiKey = "SETLISTFM_KEY"
)

type Client struct {
	accept  string
	apiKey  string
	baseUrl string
	http    *http.Client
}

func NewClient() *Client {
	return &Client{
		apiKey:  os.Getenv(ApiKey),
		accept:  "application/json",
		baseUrl: baseAddress,
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) AcceptJson() {
	c.accept = "application/json"
}

func (c *Client) AcceptXml() {
	c.accept = "application/xml"
}

func (c *Client) SetAuthInfo(apiKey string) {
	c.apiKey = apiKey
}

func (c *Client) get(url string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", c.accept)
	req.Header.Add("X-API-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// TODO: decode error in response.
		log.Fatal("HTTP NOT OK")
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

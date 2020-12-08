package setlistfm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type Error struct {
	Message string `json:"message"`
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
		return fmt.Errorf("setlistfm: HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("setlistfm: couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	return fmt.Errorf("setlistfm: %v", e)
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
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

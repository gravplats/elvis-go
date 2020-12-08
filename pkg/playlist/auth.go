package playlist

import (
	"context"
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func getClient() (*spotify.Client, error) {
	// White-listed addresses to redirect to after authentication success OR failure
	// See: https://developer.spotify.com/dashboard/
	var redirectURLPort = ":5555"
	var redirectURL = "http://localhost" + redirectURLPort + "/callback"

	auth := spotify.NewAuthenticator(redirectURL,
		spotify.ScopePlaylistModifyPrivate,
		// For the "from_token" option in market.
		// https://developer.spotify.com/documentation/web-api/reference/search/search/
		spotify.ScopeUserReadPrivate,
	)

	token, err := getToken(redirectURLPort, &auth)
	if err != nil {
		return nil, err
	}

	client := auth.NewClient(token)

	return &client, nil
}

func getToken(addr string, auth *spotify.Authenticator) (*oauth2.Token, error) {
	var token *oauth2.Token

	filename, err := getTokenCacheFilename()
	if err != nil {
		debug.Println(err)
	}

	// If we have already have a token then let's use that one. This avoids the step of creating a HTTP server for the
	// callback URL, and opening the callback URL in a browser window which moves focus from the CLI app to the browser.
	if filename != "" {
		token, err := getTokenFromFile(filename)
		if err == nil {
			return token, nil
		}

		debug.Println(err)
	}

	token, err = getTokenFromWeb(addr, auth)
	if err != nil {
		return nil, err
	}

	if filename != "" {
		err = saveTokenToFile(filename, token)
		if err != nil {
			debug.Println(err)
		}
	}

	return token, nil
}

func getTokenCacheFilename() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	tokenCacheDir := filepath.Join(home, ".credentials")
	os.MkdirAll(tokenCacheDir, 0755)

	return filepath.Join(tokenCacheDir, "elvis-spotify.json"), nil
}

func getTokenFromFile(filename string) (*oauth2.Token, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var token *oauth2.Token
	err = json.Unmarshal(buf, &token)
	if err != nil {
		return nil, err
	}

	empty := oauth2.Token{}
	if empty == *token {
		return nil, err
	}

	return token, nil
}

func getTokenFromWeb(addr string, auth *spotify.Authenticator) (*oauth2.Token, error) {
	type Result struct {
		Err   error
		Token *oauth2.Token
	}

	const state = "elvis-cli"

	var ch = make(chan Result)

	handler := http.NewServeMux()
	server := http.Server{Addr: addr, Handler: handler}

	go func(s http.Server, h *http.ServeMux) {
		h.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			token, err := auth.Token(state, r)
			if err != nil {
				ch <- Result{
					Err:   err,
					Token: nil,
				}
			}

			ch <- Result{
				Err:   nil,
				Token: token,
			}

			w.Header().Set("Content-Type", "text/html")

			// Pure evil! Close the browser window. Unfortunately the browser will still have focus.
			io.WriteString(w, "<script>window.close()</script>")
		})

		err := s.ListenAndServe()
		if err != nil {
			ch <- Result{
				Err:   err,
				Token: nil,
			}
		}
	}(server, handler)

	url := auth.AuthURL(state)
	browser.OpenURL(url)

	result := <-ch

	server.Shutdown(context.Background())

	return result.Token, result.Err
}

func saveTokenToFile(filename string, token *oauth2.Token) error {
	buf, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, buf, 0644)
	if err != nil {
		return err
	}

	return nil
}

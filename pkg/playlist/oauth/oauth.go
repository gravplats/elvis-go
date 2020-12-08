package oauth

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GetToken(addr string, auth *spotify.Authenticator) *oauth2.Token {
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
			return token
		}

		debug.Println(err)
	}

	token = getTokenFromWeb(addr, auth)
	if filename != "" {
		err = saveTokenToFile(filename, token)
		if err != nil {
			debug.Println(err)
		}
	}

	return token
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

func getTokenFromWeb(addr string, auth *spotify.Authenticator) *oauth2.Token {
	const state = "elvis-cli"

	var ch = make(chan *oauth2.Token)

	go func() {
		h := http.NewServeMux()
		s := http.Server{Addr: addr, Handler: h}

		h.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			token, err := auth.Token(state, r)
			if err != nil {
				// TODO: should pass error back.
				log.Fatal(err)
			}

			ch <- token

			w.Header().Set("Content-Type", "text/html")

			// Pure evil! Close the browser window. Unfortunately the browser will still have focus.
			io.WriteString(w, "<script>window.close()</script>")
		})

		err := s.ListenAndServe()
		if err != nil {
			// TODO: should pass error back.
			log.Fatal(err)
		}

		// TODO: figure out how to best shutdown the server.
	}()

	url := auth.AuthURL(state)
	browser.OpenURL(url)

	token := <-ch

	return token
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

package oauth

import (
	"encoding/json"
	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func requestToken(addr string, auth *spotify.Authenticator) *oauth2.Token {
	const state = "elvis-cli"

	var ch = make(chan *oauth2.Token)

	go func() {
		h := http.NewServeMux()
		s := http.Server{Addr: addr, Handler: h}

		h.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			token, err := auth.Token(state, r)
			if err != nil {
				log.Fatal(err)
			}

			ch <- token

			w.Header().Set("Content-Type", "text/html")

			// Pure evil! Close the browser window. Unfortunately the browser will still have focus.
			io.WriteString(w, "<script>window.close()</script>")
		})

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

		// TODO: figure out how to best shutdown the server.
	}()

	url := auth.AuthURL(state)
	browser.OpenURL(url)

	token := <-ch

	return token
}

func readToken(dir string) (*oauth2.Token, bool) {
	buf, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, false
	}

	var token *oauth2.Token
	err = json.Unmarshal(buf, &token)
	if err != nil {
		return nil, false
	}

	empty := oauth2.Token{}
	if empty == *token {
		return nil, false
	}

	// Let's assume that the token contains valid values here.
	return token, true
}

func GetToken(addr string, auth *spotify.Authenticator) *oauth2.Token {
	// TODO: perhaps this path should be configurable.
	var dir = "oauth-token.json"

	// If we have already have a token then let's use that one. This avoids the step of creating a HTTP server for the
	// callback URL, and opening the callback URL in a browser window which moves focus from the CLI app to the browser.
	if token, ok := readToken(dir); ok {
		return token
	}

	token := requestToken(addr, auth)

	// Save token for future needs.
	buf, err := json.Marshal(token)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(dir, buf, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	return token
}

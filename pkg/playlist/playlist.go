package playlist

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist/oauth"
	"github.com/mrydengren/elvis/pkg/spinner"
	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
)

// TODO: add New? for this new in order to set default `Type`.
type ItemGroup struct {
	Name  string
	Items []Item
	Type  Type
}

type Item struct {
	Artist string
	Name   string
}

type Type struct {
	FilterField string
	SearchType  spotify.SearchType
}

var (
	ItemGroupTypeAlbum = Type{
		FilterField: "album",
		SearchType:  spotify.SearchTypeAlbum,
	}
	ItemGroupTypeTrack = Type{
		FilterField: "track",
		SearchType:  spotify.SearchTypeTrack,
	}
)

func Create(itemGroup ItemGroup) error {
	spinner.Start("Fetching API access token from Spotify Accounts service.")

	// White-listed addresses to redirect to after authentication success OR failure
	// See: https://developer.spotify.com/dashboard/
	var redirectURLPort = ":5555"
	var redirectURL = "http://localhost" + redirectURLPort + "/callback"

	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopePlaylistModifyPrivate)
	token := oauth.GetToken(redirectURLPort, &auth)

	client := auth.NewClient(token)

	// TODO: how to handle Fail()?
	spinner.Succeed()

	spinner.Start("Fetching Spotify user information.")

	user, err := client.CurrentUser()
	if err != nil {
		spinner.Fail()
		return err
	}

	debug.DumpJson(user, "spotify-user.json")

	spinner.Succeed()

	var matches []Match
	var trackIds []spotify.ID

	spinner.Start("Fetching Spotify tracks.")

	switch itemGroup.Type.FilterField {
	case "album":
		resources := search(&client, itemGroup)
		matches = match(itemGroup, resources)

		var albumIds []spotify.ID
		for _, match := range matches {
			if match.ID == "" {
				continue
			}

			albumIds = append(albumIds, match.ID)
		}

		pages := getAlbumTracks(&client, albumIds)

		for _, page := range pages {
			for _, track := range page.Tracks {
				trackIds = append(trackIds, track.ID)
			}
		}
	case "track":
		resources := search(&client, itemGroup)
		matches = match(itemGroup, resources)

		for _, match := range matches {
			if match.ID == "" {
				continue
			}

			trackIds = append(trackIds, match.ID)
		}
	}

	// TODO: how to handle Fail()?
	spinner.Succeed()

	spinner.Start(fmt.Sprintf("Creating Spotify playlist for %s.", itemGroup.Name))

	playlist, err := client.CreatePlaylistForUser(user.ID, itemGroup.Name, "", false)
	if err != nil {
		spinner.Fail()
		return err
	}

	spinner.Succeed()

	spinner.Start("Adding Spotify tracks to playlist.")
	// It is only possible to add 100 tracks at a time to a playlist.
	for _, slice := range Split(trackIds, 100) {
		_, err = client.AddTracksToPlaylist(playlist.ID, slice...)
		if err != nil {
			spinner.Fail()
			return err
		}
	}

	spinner.Succeed()

	// Add extra newline
	fmt.Println("")

	for _, track := range matches {
		// Don't really need the spinner for this, but everything we need is already there.
		spinner.Start(fmt.Sprintf("%s - %s", track.Artist, track.Name))
		spinner.Stop(track.ID != "")
	}

	browser.OpenURL(string(playlist.URI))

	return nil
}

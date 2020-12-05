package playlist

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist/oauth"
	"github.com/mrydengren/elvis/pkg/spinner"
	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
	"log"
)

func Create(searchItemGroup SearchItemGroup) {
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
		log.Fatal(err)
	}

	debug.DumpJson(user, "spotify-user.json")

	spinner.Succeed()

	spinner.Start("Fetching Spotify tracks.")

	resources := search(&client, searchItemGroup)
	matches := match(searchItemGroup, resources)

	// TODO: how to handle Fail()?
	spinner.Succeed()

	spinner.Start(fmt.Sprintf("Creating Spotify playlist for %s.", searchItemGroup.Artist))

	playlist, err := client.CreatePlaylistForUser(user.ID, searchItemGroup.Artist, "", false)
	if err != nil {
		spinner.Fail()
		log.Fatal(err)
	}

	spinner.Succeed()

	spinner.Start("Adding Spotify track to playlist.")

	var IDs []spotify.ID
	for _, track := range matches {
		if track.ID == "" {
			continue
		}

		IDs = append(IDs, track.ID)
	}

	_, err = client.AddTracksToPlaylist(playlist.ID, IDs...)
	if err != nil {
		spinner.Fail()
		log.Fatal(err)
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
}

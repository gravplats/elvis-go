package cmd

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/cmd/parse"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
	"log"
)

func Setlist(value string) {
	setlistId := parse.SetlistId(value)
	if setlistId == "" {
		log.Fatal("Missing or incorrect <value>")
	}

	spinner.Start(fmt.Sprintf("Fetching setlist tracks for ID %s.", setlistId))

	client := setlistfm.NewClient()

	setlist, err := client.Setlist(setlistId)
	if err != nil {
		spinner.Fail()
		log.Fatal(err)
	}

	tracklist := playlist.FromSetlist(setlist)

	spinner.Succeed()

	playlist.Create(tracklist)
}
package setlist

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
	"log"
)

func Setlist(value string) {
	setlistId := SetlistId(value)
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

	debug.DumpJson(setlist, "setlistfm-setlist.json")

	searchItemGroup := FromSetlist(setlist)

	spinner.Succeed()

	playlist.Create(searchItemGroup)
}

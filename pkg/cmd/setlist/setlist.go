package setlist

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
)

func Setlist(value string) error {
	setlistId, err := parseSetlistID(value)
	if err != nil {
		return err
	}

	spinner.Start(fmt.Sprintf("Fetching setlist tracks for ID %s.", setlistId))

	client := setlistfm.NewClient()
	setlist, err := client.Setlist(setlistId)
	if err != nil {
		spinner.Fail()
		return err
	}

	debug.DumpJson(setlist, "setlistfm-setlist.json")

	searchItemGroup := FromSetlist(setlist)

	spinner.Succeed()

	return playlist.Create(searchItemGroup)
}

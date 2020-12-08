package yt

import (
	"context"
	"fmt"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func Yt(value string) error {
	youtubeId, err := parseYouTubeID(value)
	if err != nil {
		return err
	}

	spinner.Start(fmt.Sprintf("Fetching YouTube data for ID %s", youtubeId))

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_KEY")))
	if err != nil {
		return err
	}

	resp, err := service.Videos.List([]string{"snippet"}).Id(value).Do()
	if err != nil {
		spinner.Fail()
		return err
	}

	if len(resp.Items) == 0 {
		spinner.Fail()
		return fmt.Errorf("Found no matches")
	}

	item := resp.Items[0]
	debug.DumpJson(item, "youtube-video.json")

	itemGroup := fromDescription(item.Snippet.Title, item.Snippet.Description)

	spinner.Succeed()

	return playlist.Create(itemGroup)
}

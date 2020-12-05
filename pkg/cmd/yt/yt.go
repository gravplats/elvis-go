package yt

import (
	"context"
	"fmt"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
)

func Yt(value string) {
	youtubeId := YouTubeId(value)
	if youtubeId == "" {
		log.Fatal("Missing or incorrect <value>")
	}

	spinner.Start(fmt.Sprintf("Fetching YouTube data for ID %s", youtubeId))

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_KEY")))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := service.Videos.List([]string{"snippet"}).Id(value).Do()
	if err != nil {
		spinner.Fail()
		log.Fatal(err)
	}

	if len(resp.Items) == 0 {
		spinner.Fail()
		log.Fatal("Found no matches")
	}

	item := resp.Items[0]
	debug.DumpJson(item, "youtube-video.json")

	searchItemGroup := FromDescription(item.Snippet.Title, item.Snippet.Description)

	spinner.Succeed()

	playlist.Create(searchItemGroup)
}

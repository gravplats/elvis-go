package playlist

import (
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/zmb3/spotify"
	"sort"
)

func getAlbumTracks(client *spotify.Client, ids []spotify.ID) []*spotify.SimpleTrackPage {
	type Result struct {
		Index int
		Value *spotify.SimpleTrackPage
	}

	ch := make(chan Result)

	for i, v := range ids {
		go func(index int, id spotify.ID) {
			value, err := client.GetAlbumTracks(id)
			if err != nil {
				ch <- Result{
					Index: index,
					Value: nil,
				}

				return
			}

			ch <- Result{
				Index: index,
				Value: value,
			}

		}(i, v)
	}

	var results []Result

	for {
		result := <-ch
		results = append(results, result)

		if len(results) == len(ids) {
			break
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	var pages []*spotify.SimpleTrackPage
	for _, result := range results {
		pages = append(pages, result.Value)
	}

	debug.DumpJson(pages, "spotify-album-tracks.json")

	return pages
}

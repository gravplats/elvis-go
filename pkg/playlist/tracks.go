package playlist

import (
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/limit"
	"github.com/zmb3/spotify"
	"sort"
	"sync"
)

func getAlbumTracks(client *spotify.Client, ids []spotify.ID) []*spotify.SimpleTrackPage {
	type Result struct {
		Index int
		Value *spotify.SimpleTrackPage
	}

	ch := make(chan Result, len(ids))

	// Let's avoid hammering the Spotify Web API with possibly thousands of simultaneous requests by limiting the
	// number of concurrent requests and the number of requests per seconds. Numbers are chosen arbitrarily.
	cl := limit.NewConcurrency(16)
	rl := limit.NewRate(16)

	wg := sync.WaitGroup{}
	wg.Add(len(ids))

	for i, v := range ids {
		go func(index int, id spotify.ID) {
			defer cl.Release()
			defer wg.Done()

			cl.Take()
			rl.Take()

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

	wg.Wait()
	rl.Stop()

	close(ch)

	var results []Result
	for result := range ch {
		results = append(results, result)
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

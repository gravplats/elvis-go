package playlist

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/mrydengren/elvis/pkg/limit"
	"github.com/zmb3/spotify"
	"sort"
	"strings"
	"sync"
)

func search(client *spotify.Client, group ItemGroup) [][]Resource {
	type Result struct {
		Index int
		Value *spotify.SearchResult
	}

	ch := make(chan Result, len(group.Items))

	// Let's avoid hammering the Spotify Web API with possibly thousands of simultaneous requests by limiting the
	// number of concurrent requests and the number of requests per seconds. Numbers are chosen arbitrarily.
	cl := limit.NewConcurrency(16)
	rl := limit.NewRate(16)

	wg := sync.WaitGroup{}
	wg.Add(len(group.Items))

	for i, v := range group.Items {
		go func(index int, item Item) {
			defer cl.Release()
			defer wg.Done()

			cl.Take()
			rl.Take()

			query := fmt.Sprintf("artist:%s %s:%s",
				strings.ToLower(item.Artist),
				group.Type.FilterField,
				strings.ToLower(item.Name),
			)

			country := "from_token"
			limit := 3

			options := spotify.Options{
				Country: &country,
				Limit:   &limit,
			}

			value, err := client.SearchOpt(query, group.Type.SearchType, &options)
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

	sort.Slice(results, func(i int, j int) bool {
		return results[i].Index < results[j].Index
	})

	var resources = make([][]Resource, 0, len(results))
	var searchResults []*spotify.SearchResult

	for _, result := range results {
		switch group.Type.FilterField {
		case "album":
			resources = append(resources, fromAlbum(result.Value))
		case "track":
			resources = append(resources, fromTrack(result.Value))
		}

		searchResults = append(searchResults, result.Value)
	}

	debug.DumpJson(searchResults, "spotify-search-results.json")

	return resources
}

package playlist

import (
	"fmt"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/zmb3/spotify"
	"log"
	"sort"
	"strings"
)

func search(client *spotify.Client, group SearchItemGroup) [][]Resource {
	type Result struct {
		Index int
		Value *spotify.SearchResult
	}

	ch := make(chan Result)

	for i, v := range group.Items {
		go func(index int, item SearchItem) {
			query := fmt.Sprintf("artist:%s %s:%s",
				strings.ToLower(item.Artist),
				group.Type.FilterField,
				strings.ToLower(item.Name),
			)

			// TODO: these values should probably be configurable.
			country := "SE"
			limit := 3

			options := spotify.Options{
				Country: &country,
				Limit:   &limit,
			}

			value, err := client.SearchOpt(query, group.Type.Search, &options)
			if err != nil {
				log.Println(err)

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

	// TODO: improve error handling, avoid sending too many concurrent requests, timeouts etc.

	for {
		result := <-ch
		results = append(results, result)

		if len(results) == len(group.Items) {
			break
		}
	}

	sort.Slice(results, func(i int, j int) bool {
		return results[i].Index < results[j].Index
	})

	var resources = make([][]Resource, 0, len(results))
	var searchResults []*spotify.SearchResult

	for _, result := range results {
		switch group.Type.FilterField {
		case "track":
			resources = append(resources, FromTrack(result.Value))
		}

		searchResults = append(searchResults, result.Value)
	}

	debug.DumpJson(searchResults, "spotify-search-results.json")

	return resources
}

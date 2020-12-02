package playlist

import (
	"fmt"
	"github.com/zmb3/spotify"
	"log"
	"sort"
	"strings"
)

func searchByTracks(client *spotify.Client, list Tracklist) []*spotify.SearchResult {
	type result struct {
		Index int
		Value *spotify.SearchResult
	}

	ch := make(chan result)

	for i, v := range list.Tracks {
		go func(index int, track Track) {
			query := fmt.Sprintf("artist:%s track:%s",
				strings.ToLower(track.Artist),
				strings.ToLower(track.Name),
			)

			// TODO: these values should probably be configurable.
			country := "SE"
			limit := 3

			options := spotify.Options{
				Country: &country,
				Limit:   &limit,
			}

			value, err := client.SearchOpt(query, spotify.SearchTypeTrack, &options)
			if err != nil {
				log.Println(err)

				ch <- result{
					Index: index,
					Value: nil,
				}

				return
			}

			// TODO: debug JSON if debug flag is set.

			ch <- result{
				Index: index,
				Value: value,
			}
		}(i, v)
	}

	var results []result

	// TODO: improve error handling, avoid sending too many concurrent requests, timeouts etc.

	for {
		result := <-ch
		results = append(results, result)

		if len(results) == len(list.Tracks) {
			break
		}
	}

	sort.Slice(results, func(i int, j int) bool {
		return results[i].Index < results[j].Index
	})

	var searchResults []*spotify.SearchResult
	for _, result := range results {
		searchResults = append(searchResults, result.Value)
	}

	return searchResults
}

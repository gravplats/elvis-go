package playlist

import (
	"github.com/agext/levenshtein"
	"github.com/zmb3/spotify"
)

type SearchType struct {
	FilterField string
	Search      spotify.SearchType
}

var (
	SearchItemTypeTrack = SearchType{
		FilterField: "track",
		Search:      spotify.SearchTypeTrack,
	}
)

type SearchItem struct {
	Artist string
	Name   string
}

// TODO: add New? for this new in order to set default `Type`.
type SearchItemGroup struct {
	Artist string
	Items  []SearchItem
	Type   SearchType
}

type SearchMatch struct {
	Artist string
	// Will be an empty string if no match was found
	ID   spotify.ID
	Name string
}

func getBestMatch(item SearchItem, resources []Resource) Resource {
	const Threshold = 0.4

	bestSim := float64(0)
	bestMatch := Resource{
		ID:   "",
		Name: "",
	}

	for _, resource := range resources {
		sim := levenshtein.Similarity(resource.Name, item.Name, levenshtein.NewParams())

		if sim == 1 {
			bestMatch = resource
			break
		}

		if sim > bestSim && sim > Threshold {
			bestSim = sim
			bestMatch = resource
		}
	}

	return bestMatch
}

func match(group SearchItemGroup, resources [][]Resource) []SearchMatch {
	var matches []SearchMatch
	for i, item := range group.Items {
		bestMatch := getBestMatch(item, resources[i])
		matches = append(matches, SearchMatch{
			Artist: item.Artist,
			ID:     bestMatch.ID,
			Name:   item.Name,
		})
	}

	return matches
}

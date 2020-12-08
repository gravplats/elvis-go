package playlist

import (
	"github.com/agext/levenshtein"
	"github.com/zmb3/spotify"
)

type Match struct {
	Artist string
	// Will be an empty string if no match was found
	ID   spotify.ID
	Name string
}

func getBestMatch(item Item, resources []Resource) Resource {
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

func match(group ItemGroup, resources [][]Resource) []Match {
	var matches []Match
	for i, item := range group.Items {
		bestMatch := getBestMatch(item, resources[i])
		matches = append(matches, Match{
			Artist: item.Artist,
			ID:     bestMatch.ID,
			Name:   item.Name,
		})
	}

	return matches
}

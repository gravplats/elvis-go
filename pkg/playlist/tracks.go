package playlist

import (
	"github.com/agext/levenshtein"
	"github.com/zmb3/spotify"
)

type Track struct {
	Artist string
	Name   string
}

type Tracklist struct {
	Artist string
	Id     string
	Tracks []Track
}

type MatchedTrack struct {
	Artist string
	// Will be an empty string if no match was found
	ID   spotify.ID
	Name string
}

func match(tracklist Tracklist, searchResults []*spotify.SearchResult) []MatchedTrack {
	var matches []MatchedTrack
	for i, track := range tracklist.Tracks {
		// TODO: perhaps this value should be configurable.
		const Threshold = 0.4

		var fullTrack spotify.FullTrack

		best := float64(0)
		found := false

		for _, v := range searchResults[i].Tracks.Tracks {
			p := levenshtein.NewParams()
			s := levenshtein.Similarity(v.Name, track.Name, p)

			if s == 1 {
				found = true
				fullTrack = v
				break
			}

			if s > best && s > Threshold {
				found = true
				fullTrack = v
			}
		}

		id := spotify.ID("")
		if found {
			id = fullTrack.ID
		}

		matches = append(matches, MatchedTrack{
			Artist: track.Artist,
			ID:     id,
			Name:   track.Name,
		})
	}

	return matches
}

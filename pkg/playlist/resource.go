package playlist

import "github.com/zmb3/spotify"

type Resource struct {
	ID   spotify.ID
	Name string
}

func fromAlbum(result *spotify.SearchResult) []Resource {
	if result == nil {
		return []Resource{}
	}

	var resources = make([]Resource, 0, len(result.Albums.Albums))
	for _, item := range result.Albums.Albums {
		resources = append(resources, Resource{
			ID:   item.ID,
			Name: item.Name,
		})
	}
	return resources
}

func fromTrack(result *spotify.SearchResult) []Resource {
	if result == nil {
		return []Resource{}
	}

	var resources = make([]Resource, 0, len(result.Tracks.Tracks))
	for _, item := range result.Tracks.Tracks {
		resources = append(resources, Resource{
			ID:   item.ID,
			Name: item.Name,
		})
	}
	return resources
}

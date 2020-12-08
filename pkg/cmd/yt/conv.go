package yt

import (
	"github.com/mrydengren/elvis/pkg/playlist"
	"regexp"
	"strings"
)

func fromDescription(title, description string) playlist.ItemGroup {
	var items []playlist.Item

	lines := strings.Split(description, "\n")

	var memento []string

	for _, line := range lines {
		if isEmpty(line) {
			memento = nil
			continue
		}

		if endsInColon(line) {
			continue
		}

		if startsWithHttpProtocol(line) {
			switch len(memento) {
			case 1:
				tokens := strings.Split(memento[0], "- ")
				items = append(items, playlist.Item{
					Artist: tokens[0],
					Name:   tokens[1],
				})
			case 2:
				items = append(items, playlist.Item{
					Artist: memento[0],
					Name:   memento[1],
				})
			default:
				// ???
			}

			memento = nil
			continue
		}

		memento = append(memento, line)
	}

	return playlist.ItemGroup{
		Name:  title,
		Items: items,
		Type:  playlist.ItemGroupTypeAlbum,
	}
}

func endsInColon(line string) bool {
	return regexp.MustCompile(`:\s*$`).MatchString(line)
}

func isEmpty(line string) bool {
	return regexp.MustCompile(`^\s*$`).MatchString(line)
}

func startsWithHttpProtocol(line string) bool {
	return regexp.MustCompile(`^https?://`).MatchString(line)
}

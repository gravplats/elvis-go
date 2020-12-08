package yt

import (
	"fmt"
	"regexp"
)

func parseYouTubeID(value string) (string, error) {
	var pattern *regexp.Regexp
	var matches []string

	pattern = regexp.MustCompile(`^.*\?v=([0-9A-Za-z]+)&?`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1], nil
	}

	pattern = regexp.MustCompile(`^([0-9A-Za-z]+)$`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1], nil
	}

	return "", fmt.Errorf("Missing or incorrect YoutTube ID")
}

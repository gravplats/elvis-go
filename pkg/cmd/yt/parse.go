package yt

import "regexp"

func YouTubeId(value string) string {
	var pattern *regexp.Regexp
	var matches []string

	pattern = regexp.MustCompile(`^.*\?v=([0-9A-Za-z]+)&?`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1]
	}

	pattern = regexp.MustCompile(`^([0-9A-Za-z]+)$`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

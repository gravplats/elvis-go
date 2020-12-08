package setlist

import (
	"fmt"
	"regexp"
)

func parseSetlistID(value string) (string, error) {
	var pattern *regexp.Regexp
	var matches []string

	pattern = regexp.MustCompile(`^.*-([0-9a-z]+)\.html$`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1], nil
	}

	pattern = regexp.MustCompile(`^([0-9a-z]+)$`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1], nil
	}

	return "", fmt.Errorf("Missing or incorrect setlist ID")
}

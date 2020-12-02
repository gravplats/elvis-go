package parse

import "regexp"

func SetlistId(value string) string {
	var pattern *regexp.Regexp
	var matches []string

	pattern = regexp.MustCompile(`^.*-([0-9a-z]+)\.html$`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1]
	}

	pattern = regexp.MustCompile(`^([0-9a-z]+)$`)
	matches = pattern.FindStringSubmatch(value)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

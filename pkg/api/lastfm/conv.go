package lastfm

func parseBool(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

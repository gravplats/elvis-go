package playlist

import "github.com/zmb3/spotify"

func MathMinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Split(values []spotify.ID, maxSize int) [][]spotify.ID {
	start := 0

	var result [][]spotify.ID

	for {
		left := len(values) - start
		take := MathMinInt(maxSize, left)

		slice := values[start:(start + take)]
		result = append(result, slice)

		if left <= maxSize {
			break
		}

		start += maxSize
	}

	return result
}

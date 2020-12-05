package yt

import "testing"

type fixture struct {
	name  string
	input string
	want  string
}

func TestParseYouTubeId(t *testing.T) {
	fixtures := []fixture{
		fixture{
			name:  "id",
			input: "aEdho6H0hFY",
			want:  "aEdho6H0hFY",
		},
		fixture{
			name:  "URL",
			input: "https://www.youtube.com/watch?v=aEdho6H0hFY",
			want:  "aEdho6H0hFY",
		},
		fixture{
			name:  "URL with timestamp",
			input: "https://www.youtube.com/watch?v=aEdho6H0hFY&t=100",
			want:  "aEdho6H0hFY",
		},
		fixture{
			name:  "empty",
			input: "",
			want:  "",
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got := YouTubeId(fixture.input)

			if got != fixture.want {
				t.Errorf("got: %+v, want: %+v", got, fixture.want)
			}
		})
	}
}

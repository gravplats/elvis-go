package yt

import "testing"

func TestParseYouTubeID(t *testing.T) {
	type fixture struct {
		name        string
		input       string
		want        string
		expectError bool
	}

	fixtures := []fixture{
		fixture{
			name:        "id",
			input:       "aEdho6H0hFY",
			want:        "aEdho6H0hFY",
			expectError: false,
		},
		fixture{
			name:        "URL",
			input:       "https://www.youtube.com/watch?v=aEdho6H0hFY",
			want:        "aEdho6H0hFY",
			expectError: false,
		},
		fixture{
			name:        "URL with timestamp",
			input:       "https://www.youtube.com/watch?v=aEdho6H0hFY&t=100",
			want:        "aEdho6H0hFY",
			expectError: false,
		},
		fixture{
			name:        "empty",
			input:       "",
			want:        "",
			expectError: true,
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got, err := parseYouTubeID(fixture.input)
			if got != fixture.want {
				t.Errorf("got: %+v, want: %+v", got, fixture.want)
			}

			if !fixture.expectError && err != nil {
				t.Errorf("Did not expected error: %+v", err)
			}

			if fixture.expectError && err == nil {
				t.Errorf("Expected error")
			}
		})
	}
}

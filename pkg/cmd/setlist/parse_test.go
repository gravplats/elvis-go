package setlist

import "testing"

func TestParseSetlistID(t *testing.T) {
	type fixture struct {
		name        string
		input       string
		want        string
		expectError bool
	}

	fixtures := []fixture{
		fixture{
			name:        "id",
			input:       "53e3dba9",
			want:        "53e3dba9",
			expectError: false,
		},
		fixture{
			name:        "URL",
			input:       "https://www.setlist.fm/setlist/opeth/2017/finlandia-talo-helsinki-finland-53e3dba9.html",
			want:        "53e3dba9",
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
			got, err := parseSetlistID(fixture.input)
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

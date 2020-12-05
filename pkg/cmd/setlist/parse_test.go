package setlist

import "testing"

func TestSetlistId(t *testing.T) {
	type fixture struct {
		name  string
		input string
		want  string
	}

	fixtures := []fixture{
		fixture{
			name:  "id",
			input: "53e3dba9",
			want:  "53e3dba9",
		},
		fixture{
			name:  "URL",
			input: "https://www.setlist.fm/setlist/opeth/2017/finlandia-talo-helsinki-finland-53e3dba9.html",
			want:  "53e3dba9",
		},
		fixture{
			name:  "empty",
			input: "",
			want:  "",
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got := SetlistId(fixture.input)

			if got != fixture.want {
				t.Errorf("got: %+v, want: %+v", got, fixture.want)
			}
		})
	}
}

package contamination_1_string

import "testing"

func TestContamination(t *testing.T) {
	cases := []struct {
		name             string
		text, char, want string
	}{
		{
			name: "unicode symbols",
			text: "привет",
			char: "z",
			want: "zzzzzz",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Contamination(tc.text, tc.char)
			if got != tc.want {
				t.Errorf("got %q want %q", got, tc.want)
			}
		})
	}
}

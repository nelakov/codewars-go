package the_feast_of_many_beasts

import "testing"

func TestFeast(t *testing.T) {
	cases := []struct {
		name        string
		beast, dish string
		want        bool
	}{
		{
			name:  "the first letters not equal",
			beast: "dog",
			dish:  "grog",
			want:  false,
		},
		{
			name:  "the first and last letters not equal",
			beast: "cat",
			dish:  "fish",
			want:  false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Feast(tc.beast, tc.dish)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

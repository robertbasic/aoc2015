package main

import "testing"

var directiontests = []struct {
	in  string
	out int
}{
	{">", 2},
	{"<", 2},
	{"^", 2},
	{"v", 2},
	{"^>v<", 4},
	{"^v^v^v^v^v", 2},
	{"^^>>vv<^", 9},
	{"vv>><>^^<<vv", 8},
}

func TestUniqueHouses(t *testing.T) {
	for _, tt := range directiontests {
		uh := UniqueHouses(tt.in)

		if uh != tt.out {
			t.Errorf("Got %d for %s, expected %d", uh, tt.in, tt.out)
		}
	}
}

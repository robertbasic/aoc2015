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

var robotdirectiontests = []struct {
	in  string
	out int
}{
	{">", 2},
	{"<", 2},
	{"^", 2},
	{"v", 2},
	{"^v", 3},
	{"v^", 3},
	{"<>", 3},
	{"><", 3},
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
	{"^^>>vv<^", 4},
	{"vv>><>^^<<vv", 7},
}

func TestUniqueHouses(t *testing.T) {
	for _, tt := range directiontests {
		uh := UniqueHouses(tt.in)

		if uh != tt.out {
			t.Errorf("Got %d for %s, expected %d", uh, tt.in, tt.out)
		}
	}
}
func TestRobotUniqueHouses(t *testing.T) {
	for _, tt := range robotdirectiontests {
		uh := RobotUniqueHouses(tt.in)

		if uh != tt.out {
			t.Errorf("Got %d for %s, expected %d", uh, tt.in, tt.out)
		}
	}
}

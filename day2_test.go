package main

import "testing"

var packagestests = []struct {
	in  string
	out int
}{
	{"2x4", 0},
	{"2x3x4", 58},
	{"1x1x10", 43},
	{"1x1x1", 7},
	{"3x2x4", 58},
	{"6x10x2", 196},
}

var ribbontests = []struct {
	in  string
	out int
}{
	{"2x4", 0},
	{"2x3x4", 34},
	{"1x1x10", 14},
	{"1x1x1", 5},
	{"3x2x4", 34},
	{"6x10x2", 136},
}

func TestOrders(t *testing.T) {
	for _, tt := range packagestests {
		wp := WrappingPaper(tt.in)

		if wp != tt.out {
			t.Errorf("Got %d for %s, expected %d", wp, tt.in, tt.out)
		}
	}
}

func TestRibbons(t *testing.T) {
	for _, tt := range ribbontests {
		o := Ribbon(tt.in)

		if o != tt.out {
			t.Errorf("Got %d for %s, expected %d", o, tt.in, tt.out)
		}
	}
}

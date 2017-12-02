package main

import "testing"

var ordertests = []struct {
	in  string
	out int64
}{
	{"2x4", 0},
	{"2x3x4", 58},
	{"1x1x10", 43},
	{"1x1x1", 7},
	{"3x2x4", 58},
	{"6x10x2", 196},
}

func TestOrders(t *testing.T) {
	for _, tt := range ordertests {
		o, _ := order(tt.in)

		if o != tt.out {
			t.Errorf("Got %d for %s, expected %d", o, tt.in, tt.out)
		}
	}
}

package main

import (
	"testing"
)

var floortests = []struct {
	in  string
	out int
}{
	{"(())", 0},
	{"()()", 0},
	{"(((", 3},
	{"(()(()(", 3},
	{"))(((((", 3},
	{"())", -1},
	{"))(", -1},
	{")))", -3},
	{")())())", -3},
}

var basementtests = []struct {
	in  string
	out int
}{
	{"(())", 0},
	{"()()", 0},
	{"(((", 0},
	{"(()(()(", 0},
	{"))(((((", 1},
	{"())", 3},
	{"))(", 1},
	{")))", 1},
	{")())())", 1},
}

func TestFloors(t *testing.T) {
	for _, tt := range floortests {
		cf := CurrentFloor(tt.in)

		if cf != tt.out {
			t.Errorf("Got %d for %s, expected %d", cf, tt.in, tt.out)
		}
	}
}

func TestBasement(t *testing.T) {
	for _, tt := range basementtests {
		bap := EnterBasementAt(tt.in)

		if bap != tt.out {
			t.Errorf("Got %d for %s, expected %d", bap, tt.in, tt.out)
		}
	}
}

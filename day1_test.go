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
		f, _ := Floor(tt.in)

		if f != tt.out {
			t.Errorf("Got %d, expected %d", f, tt.out)
		}
	}
}

func TestBasement(t *testing.T) {
	for _, tt := range basementtests {
		f, _ := Basement(tt.in)

		if f != tt.out {
			t.Errorf("Got %d for %s, expected %d", f, tt.in, tt.out)
		}
	}
}

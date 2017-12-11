package main

import "testing"

var nicewordtests = []struct {
	in  string
	out bool
}{
	{"ugknbfddgicrmopn", true},
	{"aaa", true},
	{"jchzalrnumimnmhp", false},
	{"haegwjzuvuyypxyu", false},
	{"dvszwmarrgswjxmb", false},
}

var vowelstests = []struct {
	in  string
	out bool
}{
	{"abc", false},
	{"aba", false},
	{"abecid", true},
	{"axxxozzzu", true},
	{"aeiou", true},
}

var doublestests = []struct {
	in  string
	out bool
}{
	{"abc", false},
	{"aba", false},
	{"abecid", false},
	{"axxxozzzu", true},
	{"aeiou", false},
	{"aaa", true},
	{"abccd", true},
}

var forbiddentests = []struct {
	in  string
	out bool
}{
	{"abc", true},
	{"aba", true},
	{"aebcid", false},
	{"axxxozzzu", false},
	{"thiscd", true},
	{"aaa", false},
	{"acxy", true},
	{"acpqdxay", true},
}

var pairtests = []struct {
	in  string
	out bool
}{
	{"xyxy", true},
	{"aaa", false},
	{"caxyzcabyz", true},
	{"aaaxyzaaa", true},
}

var repeattests = []struct {
	in  string
	out bool
}{
	{"axa", true},
	{"aaxaa", true},
	{"aax", false},
	{"aaxxaa", false},
}

var verynicewordtests = []struct {
	in  string
	out bool
}{
	{"qjhvhtzxzqqjkmpb", true},
	{"xxyxx", true},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
}

func TestIsNiceWord(t *testing.T) {
	for _, tt := range nicewordtests {
		r := IsNiceWord(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}

func TestIsVeryNiceWord(t *testing.T) {
	for _, tt := range verynicewordtests {
		r := IsVeryNiceWord(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}

func TestHasThreeVowels(t *testing.T) {
	for _, tt := range vowelstests {
		r := HasThreeVowels(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}
func TestHasDoubleLetter(t *testing.T) {
	for _, tt := range doublestests {
		r := HasDoubleLetter(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}
func TestHasForbiddenSubstring(t *testing.T) {
	for _, tt := range forbiddentests {
		r := HasForbiddenSubstring(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}

func TestHasTwoPairLetters(t *testing.T) {
	for _, tt := range pairtests {
		r := HasTwoPairLetters(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}

func TestHasRepeatedLetter(t *testing.T) {
	for _, tt := range repeattests {
		r := HasRepeatedLetter(tt.in)

		if r != tt.out {
			t.Errorf("Got %t for %s, expected %t", r, tt.in, tt.out)
		}
	}
}

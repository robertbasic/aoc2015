package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Day5 solves day 5 puzzles
func Day5() {
	var nw int
	var vnw int

	file, _ := os.Open("./inputs/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if IsNiceWord(line) {
			nw++
		}

		if IsVeryNiceWord(line) {
			vnw++
		}
	}

	fmt.Println("Number of nice words: ", nw)
	fmt.Println("Number of very nice words: ", vnw)
}

// IsVeryNiceWord checks is s a very nice string
func IsVeryNiceWord(s string) bool {
	return HasTwoPairLetters(s) && HasRepeatedLetter(s)
}

// HasTwoPairLetters checks if s has a pair of any two letters
// without overlapping
func HasTwoPairLetters(s string) bool {
	for i, j := 0, 1; i < len(s); i, j = i+1, j+1 {
		if j == len(s) {
			continue
		}

		sub := string(s[i]) + string(s[j])

		if strings.Count(s, sub) >= 2 {
			return true
		}
	}

	return false
}

// HasRepeatedLetter checks if s has a repeated letter exactly
// one letter apart
func HasRepeatedLetter(s string) bool {
	var j int
	for i := 0; i < len(s); i++ {
		j = i + 2
		if j >= len(s) {
			continue
		}

		if s[i] == s[j] {
			return true
		}
	}
	return false
}

// IsNiceWord checks is s a nice string
func IsNiceWord(s string) bool {
	return HasThreeVowels(s) && HasDoubleLetter(s) && !HasForbiddenSubstring(s)
}

// HasThreeVowels checks if s has 3 vowels
func HasThreeVowels(s string) bool {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	seen := []string{}
	chs := strings.Split(s, "")

	for _, ch := range chs {
		if vowels[ch] {
			seen = append(seen, ch)
		}
	}

	return len(seen) >= 3
}

// HasDoubleLetter checks if s has double letters
func HasDoubleLetter(s string) bool {
	for i := 97; i <= 122; i++ {
		c := string(i)
		if strings.Contains(s, fmt.Sprintf("%s%s", c, c)) {
			return true
		}
	}
	return false
}

// HasForbiddenSubstring checks if s has forbidden substrings
func HasForbiddenSubstring(s string) bool {
	fs := map[string]bool{
		"ab": true,
		"cd": true,
		"pq": true,
		"xy": true,
	}

	for f := range fs {
		if strings.Contains(s, f) {
			return true
		}
	}

	return false
}

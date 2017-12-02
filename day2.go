package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Day2 solves the day 2 puzzles
func Day2() {
	var twp int
	var tr int

	file, _ := os.Open("./inputs/day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		wp := WrappingPaper(line)
		r := Ribbon(line)
		twp = twp + wp
		tr = tr + r
	}

	fmt.Println("Need wrapping paper: ", twp)
	fmt.Println("Need ribbon length: ", tr)
}

// WrappingPaper calculates the wrapping
// paper require to wrap one present
func WrappingPaper(dimensions string) int {
	var twp int

	ns, err := Numbers(dimensions)

	if err != nil {
		return twp
	}

	dims := Dimensions(ns)

	sides := Sides(dims[0], dims[1], dims[2])

	ss := sides[0]

	twp = sides[0]*2 + sides[1]*2 + sides[2]*2 + ss

	return twp
}

// Ribbon calculates the ribbon length
// required to put a ribbon on one present
func Ribbon(dimensions string) int {
	var length int

	ns, err := Numbers(dimensions)

	if err != nil {
		return length
	}

	dims := Dimensions(ns)

	ribbon := dims[0]*2 + dims[1]*2
	bow := dims[0] * dims[1] * dims[2]

	length = ribbon + bow

	return length
}

// Numbers splits a line representing one package
// dimensions by the "x"
func Numbers(dimensions string) ([]string, error) {
	ns := strings.Split(dimensions, "x")

	if len(ns) != 3 {
		return nil, errors.New("Incorrect dimensions")
	}

	return ns, nil
}

// Dimensions converts an array of strings
// into integers that represent the actual
// dimensions and sorts that array of ints
func Dimensions(numbers []string) [3]int {
	var nrs [3]int

	for i, n := range numbers {
		nrs[i], _ = strconv.Atoi(n)
	}

	sort.Ints(nrs[:])

	return nrs
}

// Sides calculates the areas of the 3
// different sides of a box
// Sides are sorted
func Sides(l int, w int, h int) [3]int {
	var sides [3]int

	sides[0] = l * w
	sides[1] = l * h
	sides[2] = h * w

	sort.Ints(sides[:])

	return sides
}

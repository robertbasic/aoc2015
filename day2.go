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

func Day2() {
	var tt int
	var tr int

	file, _ := os.Open("./inputs/day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		o, _ := Order(line)
		r := Ribbon(line)
		tt = tt + o
		tr = tr + r
	}

	fmt.Println("Order a total of: ", tt)
	fmt.Println("Need ribbon length: ", tr)
}

func Order(dimensions string) (int, error) {
	var t int

	ns, err := Numbers(dimensions)

	if err != nil {
		return t, err
	}

	l, _ := strconv.Atoi(ns[0])
	w, _ := strconv.Atoi(ns[1])
	h, _ := strconv.Atoi(ns[2])

	sides := Sides(l, w, h)

	s := Smallest(sides)

	t = sides[0]*2 + sides[1]*2 + sides[2]*2 + s

	return t, nil
}

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

func Numbers(dimensions string) ([]string, error) {
	ns := strings.Split(dimensions, "x")

	if len(ns) != 3 {
		return nil, errors.New("Incorrect dimensions")
	}

	return ns, nil
}

func Dimensions(numbers []string) [3]int {
	l, _ := strconv.Atoi(numbers[0])
	w, _ := strconv.Atoi(numbers[1])
	h, _ := strconv.Atoi(numbers[2])

	nrs := [3]int{l, w, h}

	sort.Ints(nrs[:])

	return nrs
}

func Sides(l int, w int, h int) [3]int {
	var sides [3]int

	sides[0] = l * w
	sides[1] = l * h
	sides[2] = h * w

	return sides
}

func Smallest(sides [3]int) int {
	s := sides[0]

	for _, a := range sides {
		if a < s {
			s = a
		}
	}

	return s
}

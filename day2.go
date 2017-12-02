package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2() {
	var tt int64

	file, _ := os.Open("./inputs/day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		o, _ := order(scanner.Text())
		tt = tt + o
	}

	fmt.Println("Order a total of: ", tt)
}

func order(dimensions string) (int64, error) {
	var t int64
	ns := strings.Split(dimensions, "x")

	if len(ns) != 3 {
		return t, errors.New("Incorrect dimensions")
	}

	l, _ := strconv.ParseInt(string(ns[0]), 10, 32)
	w, _ := strconv.ParseInt(string(ns[1]), 10, 32)
	h, _ := strconv.ParseInt(string(ns[2]), 10, 32)

	sides := sides(l, w, h)

	s := smallest(sides)

	t = sides[0]*2 + sides[1]*2 + sides[2]*2 + s

	return t, nil
}

func sides(l int64, w int64, h int64) [3]int64 {
	var sides [3]int64

	sides[0] = l * w
	sides[1] = l * h
	sides[2] = h * w

	return sides
}

func smallest(sides [3]int64) int64 {
	s := sides[0]

	for _, a := range sides {
		if a < s {
			s = a
		}
	}

	return s
}

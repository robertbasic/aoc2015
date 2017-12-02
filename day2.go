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
	var tt int

	file, _ := os.Open("./inputs/day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		o, _ := Order(scanner.Text())
		tt = tt + o
	}

	fmt.Println("Order a total of: ", tt)
}

func Order(dimensions string) (int, error) {
	var t int

	ns, err := numbers(dimensions)

	if err != nil {
		return t, err
	}

	l, _ := strconv.Atoi(ns[0])
	w, _ := strconv.Atoi(ns[1])
	h, _ := strconv.Atoi(ns[2])

	sides := sides(l, w, h)

	s := smallest(sides)

	t = sides[0]*2 + sides[1]*2 + sides[2]*2 + s

	return t, nil
}

func Ribbon(dimensions string) int {
	var length int

	return length
}

func numbers(dimensions string) ([]string, error) {
	ns := strings.Split(dimensions, "x")
	fmt.Println(ns)
	if len(ns) != 3 {
		return nil, errors.New("Incorrect dimensions")
	}

	return ns, nil
}

func sides(l int, w int, h int) [3]int {
	var sides [3]int

	sides[0] = l * w
	sides[1] = l * h
	sides[2] = h * w

	return sides
}

func smallest(sides [3]int) int {
	s := sides[0]

	for _, a := range sides {
		if a < s {
			s = a
		}
	}

	return s
}

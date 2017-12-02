package main

import (
	"bufio"
	"fmt"
	"os"
)

// Position holds the current
// position
type Position struct {
	X int
	Y int
}

func (p *Position) move(x int, y int) {
	p.X = p.X + x
	p.Y = p.Y + y
}
func (p Position) coords() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

// Day3 solves day 3 puzzles
func Day3() {
	var uh int

	file, _ := os.Open("./inputs/day3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		uh = uh + UniqueHouses(line)
	}

	fmt.Println("Unique houses visited: ", uh)
}

// UniqueHouses finds the number
// of unique houses packages were
// delivered to
func UniqueHouses(directions string) int {
	var uh int

	cp := Position{}

	houses := make(map[string]int)

	houses[cp.coords()]++

	for _, dir := range directions {
		x, y := translate(dir)
		cp.move(x, y)
		houses[cp.coords()]++
	}

	uh = len(houses)

	return uh
}

func translate(direction rune) (int, int) {
	if direction == rune('^') {
		return 1, 0
	} else if direction == rune('>') {
		return 0, 1
	} else if direction == rune('v') {
		return -1, 0
	} else if direction == rune('<') {
		return 0, -1
	}

	return 0, 0
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Day1 solves the day 1 puzzles
func Day1() {
	var directions string

	file, _ := os.Open("./inputs/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		directions = scanner.Text()
	}

	cf := CurrentFloor(directions)
	bap := EnterBasementAt(directions)
	fmt.Println("The current floor is: ", cf)
	fmt.Println("Enters the basement at position: ", bap)
}

// CurrentFloor calculates the current floor
// after following all the directions
func CurrentFloor(directions string) int {
	var cf int

	for _, dir := range directions {
		calculateCurrentFloor(dir, &cf)
	}

	return cf
}

// EnterBasementAt calculates the position at which
// the basement is entered
func EnterBasementAt(directions string) int {
	var cf int
	var p int

	for i, dir := range directions {
		calculateCurrentFloor(dir, &cf)

		if cf < 0 {
			p = i + 1
			break
		}
	}

	return p
}

func calculateCurrentFloor(direction rune, cf *int) {
	if direction == rune('(') {
		*cf++
	} else if direction == rune(')') {
		*cf--
	}
}

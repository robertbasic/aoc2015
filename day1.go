package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Day1() {
	var floors string

	file, _ := os.Open("./inputs/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		floors = scanner.Text()
	}

	floor, _ := Floor(floors)
	basement, _ := Basement(floors)
	fmt.Println("The floor is: ", floor)
	fmt.Println("Enters the basement at position: ", basement)
}

func Floor(floors string) (int, error) {
	var f int

	var u = rune('(')
	var d = rune(')')

	for _, c := range floors {
		if c == u {
			f++
		} else if c == d {
			f--
		} else {
			return 0, errors.New("Unknown floor operation")
		}
	}

	return f, nil
}

func Basement(floors string) (int, error) {
	var f int
	var p int

	var u = rune('(')
	var d = rune(')')

	for i, c := range floors {
		if c == u {
			f++
		} else if c == d {
			f--
		} else {
			return 0, errors.New("Unknown floor operation")
		}

		if f < 0 {
			p = i + 1
			break
		}
	}

	return p, nil
}

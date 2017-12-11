package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Light represents a light
type Light struct {
	brightness int
	state      bool
	x          int
	y          int
}

func (l *Light) do(op string) {
	switch op {
	case "on":
		l.on()
	case "off":
		l.off()
	case "toggle":
		l.toggle()
	}
}

func (l *Light) dob(op string) {
	switch op {
	case "on":
		l.incb(1)
	case "off":
		l.decb(1)
	case "toggle":
		l.incb(2)
	}
}

func (l *Light) incb(b int) {
	l.brightness += b
}

func (l *Light) decb(b int) {
	l.brightness -= b
	if l.brightness < 0 {
		l.brightness = 0
	}
}

func (l *Light) on() {
	l.state = true
}
func (l *Light) off() {
	l.state = false
}
func (l *Light) toggle() {
	l.state = !l.state
}
func (l Light) coord() string {
	return fmt.Sprintf("%d-%d", l.x, l.y)
}

// Day6 solves the puzzles for day 6
func Day6() {
	file, _ := os.Open("./inputs/day6.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lights := BuildLights(1000)

	for scanner.Scan() {
		line := scanner.Text()
		lights = SetupLights(line, lights)
	}

	c := CountOnLights(lights)
	b := CountBrightness(lights)

	fmt.Println("Total lights on: ", c)
	fmt.Println("Total brightness: ", b)
}

// CountOnLights counts lights that are on
func CountOnLights(lights map[string]Light) int {
	var c int

	for _, l := range lights {
		if l.state == true {
			c++
		}
	}

	return c
}

// CountBrightness counts the total brightness
func CountBrightness(lights map[string]Light) int {
	var b int

	for _, l := range lights {
		b += l.brightness
	}

	return b
}

// SetupLights sets up lights according to instruction
func SetupLights(instruction string, lights map[string]Light) map[string]Light {
	instructions := strings.Fields(instruction)
	var action string
	var istart, jstart, iend, jend int

	if len(instructions) == 5 {
		action = instructions[1]
		istart, jstart = parseCoord(instructions[2])
		iend, jend = parseCoord(instructions[4])
	} else {
		action = instructions[0]
		istart, jstart = parseCoord(instructions[1])
		iend, jend = parseCoord(instructions[3])
	}

	for i := istart; i <= iend; i++ {
		for j := jstart; j <= jend; j++ {
			k := fmt.Sprintf("%d-%d", i, j)
			l := lights[k]
			l.do(action)
			l.dob(action)
			lights[k] = l
		}
	}

	return lights
}

func parseCoord(coord string) (int, int) {
	coords := strings.Split(coord, ",")

	s, _ := strconv.Atoi(coords[0])
	e, _ := strconv.Atoi(coords[1])

	return s, e
}

// BuildLights creates a grid of lights of the
// given size x size
func BuildLights(size int) map[string]Light {
	var lights = make(map[string]Light)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			l := Light{x: i, y: j}
			lights[l.coord()] = l
		}
	}

	return lights
}

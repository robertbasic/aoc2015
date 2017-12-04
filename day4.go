package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Day4 solves the puzzles for day 4
func Day4() {
	i := Findmd5("yzbqklnj")
	fmt.Println("The int is: ", i)
}

// Findmd5 finds int suffix to input that creates
// the md5 that begins with 5 zeros
func Findmd5(input string) int {
	var i int
	var found bool

	for !found {
		i++

		cs := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		s := fmt.Sprintf("%x", cs)

		if strings.HasPrefix(s, "000000") {
			found = true
		}
	}

	return i
}

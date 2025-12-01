package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("example.txt")
	check(err)
	position, total := 50, 0

	datString := string(dat[:])

	scanner := bufio.NewScanner(strings.NewReader(datString))
	for scanner.Scan() {
		position, total = turnTheKnob(scanner.Text(), position, total)
	}
	fmt.Printf("Total = %#v", total)

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
}

func turnTheKnob(line string, position int, total int) (int, int) {
	value, err := strconv.Atoi(line[1:len(line)])
	check(err)
	ex2 := true
	// increase := false

	if line[0] == 'L' {
		position = position - value
	} else if line[0] == 'R' {
		position = position + value
	}
	for position > 99 {
		position = position - 100
		if ex2 {
			total += 1
		}
	}
	for position < 0 {
		position = position + 100
		if ex2 {
			total += 1
		}
	}
	if position == 0 {
		// total += 1
	}

	// if increase {
	// 	// total += 1
	// }
	fmt.Printf("The dial is rotated %#v to point at %#v and total is %#v\n", line, position, total)

	return position, total
}

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
	dat, err := os.ReadFile("input.txt")
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

func turnTheKnob(line string, initialposition int, total int) (int, int) {
	value, err := strconv.Atoi(line[1:len(line)])
	check(err)
	// increase := false
	direction := 1
	if line[0] == 'L' {
		direction = -1
	}

	for value > 99 {
		total += 1
		value -= 100
	}

	position := initialposition + (direction * value)

	if position > 99 {
		position = position - 100
		if initialposition != 0 {
			total += 1
		}
	} else if position < 0 {
		position = position + 100
		if initialposition != 0 {
			total += 1
		}
	} else if position == 0 {
		total += 1
	}

	// if increase {
	// 	// total += 1
	// }
	fmt.Printf("The dial is rotated %#v to point at %#v and total is %#v\n", line, position, total)

	return position, total
}

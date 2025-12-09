package main

import (
	"fmt"
	"os"
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
	data := string(dat[:])

	values := strings.Split(data, "\n")
	var total int64 = 0
	var rollDetected bool = true

	// width := len(values[1])
	height := len(values) - 1
	var valuesNew = make([]string, height+1)

	// fmt.Printf("width is %#v and height is %#v\n", width, height)
	// fmt.Printf("total is %#v\n", total)

	// fmt.Printf("hallo %#v\n", values[0])
	for rollDetected {
		var totalRolls int64 = 0
		valuesNew, totalRolls, rollDetected = findRolls(values)
		total += totalRolls
		// fmt.Print("newvalues %#v", valuesNew)
		values = valuesNew
	}
	fmt.Printf("found %#v rolls", total)

}

func findRolls(values []string) ([]string, int64, bool) {
	width := len(values[1])
	height := len(values) - 1
	ncells := width * height
	rollDetected := false
	var valuesNew = make([]string, height+1)
	var total int64 = 0
	for i := 0; i < ncells; i++ {
		// fmt.Printf("i is %#v\n ", i)
		character := "@"
		positionx, positiony := i%width, (i-i%width)/width
		// fmt.Printf("x is %d y is %d\n", positionx, positiony)
		if string(values[positiony][positionx]) == character { // [valuesHeight, valuesWidth ]
			// fmt.Printf("found an initial %#v and at position %#v\n", character, i)
			// fmt.Printf("Scanning for more %s's\n", character)
			validRoll := charscan(values, positionx, positiony, width, height, character)
			if validRoll {
				valuesNew[positiony] = valuesNew[positiony] + "X"
				total += 1
				rollDetected = true
			} else {
				valuesNew[positiony] = valuesNew[positiony] + "@"
			}
		} else {
			valuesNew[positiony] = valuesNew[positiony] + "."
		}
	}
	// fmt.Printf("values: %#v\n", values)
	// fmt.Printf("valuesnew: %#v\n", valuesNew)
	// fmt.Printf("Total: %#v", total)
	return valuesNew, total, rollDetected
}

func charscan(values []string, x int, y int, w int, h int, char string) bool {
	total := 0
	var startx int = -1
	var endx int = 1
	var starty int = -1
	var endy int = 1
	if x == 0 {
		startx = 0
	} else {
		startx = x - 1
	}

	if x == w-1 {
		endx = w - 1
	} else {
		endx = x + 1
	}

	if y == 0 {
		starty = 0
	} else {
		starty = y - 1
	}

	if y == h-1 {
		endy = y
	} else {
		endy = y + 1
	}

	// fmt.Printf("startx %#v, endx %#v\n", startx, endx)
	// fmt.Printf("x %#v, y %#v\n", x, y)
	// fmt.Printf("starty %#v, endy %#v\n", starty, endy)
	for i := startx; i <= endx; i++ {
		// fmt.Printf("XXXX\n")
		// fmt.Printf("iii is %#v\n", i)
		for j := starty; j <= endy; j++ {
			// fmt.Printf("YYYY\n")
			// fmt.Printf("found a %#v at position x %#v and y %#v\n", string(values[j][i]), i, j)
			if string(values[j][i]) == char {
				// fmt.Printf("found for a %#v at position x %#v and y %#v\n!", char, i, j)
				if i == x && j == y {
					total -= 1
				}
				total += 1
				// fmt.Printf("total is now %#v \n!", total)
			}
		}
	}

	if total >= 4 {
		return false
	} else {
		return true
	}

}

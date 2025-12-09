package main

import (
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
	data := string(dat[:])

	values := strings.Split(data, "\n")
	var total int64 = 0

	height := len(values) - 1
	operators := strings.Fields(values[height])
	width := len(operators)
	/// pseudologic:
	/// For lines in data
	/// newValues -> append that value
	/// newValues : [1 2 3]; [4 5 6] ; [7 8 9]
	newValues := make([][]string, height)
	for i := range newValues {
		newValues[i] = make([]string, width)
	}
	for i := range newValues {
		number := strings.Fields(values[i])
		for j := 0; j < len(number); j++ {
			newValues[i][j] = number[j]
		}
	}

	total += multiply(newValues, operators)
	fmt.Printf("Total: %#v", total)

}

func multiply(values [][]string, operators []string) int64 {
	total := int64(0)
	for i := range operators {
		totalIteration := int64(0)
		for j := range values {
			if j == 0 {
				value, _ := strconv.Atoi(values[j][i])
				totalIteration += int64(value)
			} else if operators[i] == "+" {
				value, _ := strconv.Atoi(values[j][i])
				totalIteration += int64(value)
			} else if operators[i] == "*" {
				value, _ := strconv.Atoi(values[j][i])
				totalIteration *= int64(value)
			} else if operators[i] == "-" {
				value, _ := strconv.Atoi(values[j][i])
				totalIteration -= int64(value)
			}
		}
		total += totalIteration
		fmt.Printf("value %#v\n", totalIteration)
	}
	return int64(total)
}

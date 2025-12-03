package main

import (
	"fmt"
	"math"
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
	values := strings.Split(string(dat[:]), ",")
	var total int64 = 0
	// var buffer bytes.Buffer

	fmt.Println(values)
	for i := range len(values) {
		total += checkDoubles(values[i])
		// buffer.WriteString(checkDoubles(values[i]))
	}
	// fmt.Println(buffer.String())
	fmt.Println("total is %#v", total)
}

func checkDoubles(value string) int64 {

	var total int64 = 0
	values := strings.Split(string(value[:]), "-")
	start, err := strconv.ParseInt(values[0], 10, 64)
	check(err)
	end, err := strconv.ParseInt(values[1], 10, 64)
	check(err)
	for i := start; i <= end; i++ {
		doubleDetected := detectDoubles(i)
		if doubleDetected {
			// fmt.Printf("hoi %#v\n", i)
			// return (strconv.Itoa(i))
			total += i
		}
	}
	return total
}

func detectDoubles(i int64) bool {
	length := intLen(i)
	if length%2 == 0 {
		// if even, check if the digit are repeating
		return stripInt(i, length)
	} else if length%2 == 1 {
		return false // if the length is odd, then values can never be doubled
	} else {
		return false // else always false
	}
}

func intLen(i int64) int64 {
	if i == 0 {
		return 1
	}
	var count int64 = 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func stripInt(i int64, length int64) bool {
	orderOfMagnitude := math.Pow(10, float64(length)/2)
	repeater := math.Floor(float64(i) / orderOfMagnitude)
	remainder := i - int64(repeater) - int64(repeater*orderOfMagnitude)
	// fmt.Printf("%#v has order of magnitude %#v and repeats %#v, remainder is %#v\n", i, orderOfMagnitude, repeater, remainder)
	if remainder == 0 {
		return true
	} else {
		return false
	}
}

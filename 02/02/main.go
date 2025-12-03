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
	fmt.Printf("total is %#v", total)
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
			// fmt.Println("doubles detected")
			total += i
		}
	}
	return total
}

func detectDoubles(i int64) bool {
	length := intLen(i)
	doublesDetected := false

	for j := int64(2); j <= length; j++ {
		if length%j == 0 {
			// check if the numbers are divisible
			// fmt.Printf("%#v is divisible by %#v\n\n", length, j)
			doublesDetected = stripInt(i, length, j)
			if doublesDetected {
				return true
			}
		} else {
			// fmt.Printf("%#v is not divisible by %#v\n", length, j)
		}
	}
	return doublesDetected
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

func stripInt(i int64, length int64, j int64) bool {
	value := i
	orderOfMagnitude := math.Pow(10, float64(length)/float64(j))
	diminisher := float64(math.Pow(orderOfMagnitude, float64(j-1)))
	repeater := math.Floor(float64(value) / diminisher)
	// if i == 2121212121 {

	// 	fmt.Printf("%#v has %#v parts with order of magnitude %#v and repeats %#v\n", i, j, orderOfMagnitude, repeater)

	// }
	remainder := value
	for k := 0; k <= int(j); k++ {
		remainder = value - int64(repeater)*int64(diminisher)
		// if i == 2121212121 {
		// fmt.Printf("%#v minus %#v is %#v\n", value, int64(repeater)*int64(diminisher), remainder)
		// fmt.Printf("dividing %#v by  %#v\n", (float64(value - int64(repeater))), orderOfMagnitude)
		// }
		diminisher = diminisher / orderOfMagnitude
		// fmt.Printf("old value %#v and remainder %#v\n", value, remainder)
		value = remainder
		// fmt.Printf("new value %#v and remainder %#v\n", value, remainder)
		// if value < 0 {
		// 	return false
		// }
		if k == int(j) && value == 0 {
			// fmt.Printf("%#v\n", i)
			return true
		}
	}
	return false
	// if value == 282828 || value == 2828 || value == 28 {
	// 	fmt.Printf("dividing %#v by  %#v\n", (float64(i)), diminisher)
	// 	fmt.Printf("%#v minus %#v is %#v\n", i, int64(repeater)-int64(repeater*orderOfMagnitude), remainder)
	// 	fmt.Printf("%#v has %#v parts with order of magnitude %#v and repeats %#v, remainder is %#v\n", i, j, orderOfMagnitude, repeater, remainder)
	// }
	return false
}

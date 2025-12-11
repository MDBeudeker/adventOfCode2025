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
	counter := 0
	check(err)
	values := strings.Split(string(dat[:]), "\n")
	var total int64 = 0

	width := len(values[0])
	// fmt.Println(values)
	for i := range len(values) {
		startnumber, startindex := findstart(values[i][0 : width-11])
		// fmt.Printf("startnumber is %#v at location %#v\n", startnumber, startindex)
		newvalue := values[i][startindex:]

		// fmt.Printf("newvalue is %#v\n", newvalue)
		if len(newvalue) > 12 {
			// fmt.Printf("value too big!\n")
			newnewvalue := reprocess(newvalue, startnumber)
			// fmt.Printf("newvalue: %#v\n", newnewvalue)
			newtotal, _ := strconv.ParseInt(newnewvalue, 10, 64)
			// fmt.Printf("adding %#v to %#v\n", newtotal, total)
			total += int64(newtotal)
			counter += 1
			// fmt.Printf("newvalue is %#v\n", (newnewvalue))
			fmt.Print((newnewvalue))
			fmt.Print("\n")
		} else {
			fmt.Print((newvalue))
			fmt.Print("\n")
			lastvalue, _ := strconv.Atoi(newvalue)
			counter += 1
			total += int64(lastvalue)
		}
	}
	fmt.Printf("total is %#v counted %#v numbers\n", total, counter)
}

func findstart(value string) (int, int) {
	// fmt.Printf("looking in %#v\n", value)
	for j := 9; j >= 0; j-- {
		index := strings.Index(value, strconv.Itoa(j))
		if index != -1 {
			// fmt.Printf("found %#v at %#v\n", j, index)
			return j, index
		}
	}
	return 0, 0
}

// func reprocess(value string, startindex int) string {
// 	counter := startindex
// 	var newvalue string
// 	// for i := 0; i < 10; i++ {
// 	// 	index := strings.Index(value, strconv.Itoa(i))
// 	// 	if index > -1 {
// 	// 		newvaluea := value[:index]
// 	// 		newvalueb := value[index+1:]
// 	// 		newvalue := newvaluea + newvalueb
// 	// 		fmt.Printf("value: %#v, newvaluea: %#v, newvalueb %#v, newvalue: %#v\n", value, newvaluea, newvalueb, newvalue)
// 	// 		value = newvalue
// 	// 	}
// 	// 	counter += 1
// 	// 	if counter == 3 {
// 	// 		return newvalue
// 	// 	}
// 	// }
// 	return newvalue
// }

// func reprocess(value string, startindex int) string {
// 	var newvalue string = ""
// 	// fmt.Printf("chopping up string %#v with length %#v\n", value, len(value))
// 	for len(newvalue) != 12 { // final newstring HAS to have length 12
// 		for j := 1; j <= startindex; {
// 			index := 0

// 			for index != -1 {
// 				// fmt.Printf("looking for a %#v in string %#v\n", j, value)
// 				index = strings.Index(value, strconv.Itoa(j))
// 				if index == -1 {
// 					// if you cant find the index, add 1 to the thing
// 					j += 1
// 				} else {
// 					newvaluea := value[:index]
// 					newvalueb := value[index+1:]
// 					newvalue = newvaluea + newvalueb
// 					value = newvalue
// 					if len(value) == 12 {
// 						// fmt.Printf("newvalue: %#v\n", newvalue)

// 						return newvalue
// 					}
// 				}
// 			}

// 		}
// 	}
// 	// fmt.Printf("newvalue: %#v\n", newvalue)
// 	return newvalue
// }

func reprocess(value string, startindex int) string {
	// keep exactly 12 digits
	remove := len(value) - 12
	result := make([]byte, 0, len(value))

	for i := 0; i < len(value); i++ {
		c := value[i]

		for len(result) > 0 && result[len(result)-1] < c && remove > 0 {
			// remove smaller digits to make space for bigger ones
			result = result[:len(result)-1]
			remove--
		}

		result = append(result, c)
	}

	// result may be longer than 12 due to remaining removals
	return string(result[:12])
}

// 987654321111+ 811111111119 + 434234234278 +888911112111
// 987654321111+ 811111111119 + 434234234278 +888911112111

// 12313232434322321733233554252322336523622233231223312552332413252132853223222239525213933
// 954454353536
// 954454353536

// 199999999999800
// 169566718649726

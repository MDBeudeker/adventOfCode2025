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
	values := strings.Split(string(dat[:]), "\n")
	var total = 0

	fmt.Println(values)
	for i := range len(values) {
		counter := 0
		firstnumber := 0
		firstindex := -1
		secondnumber := 0
		secondindex := -1
		leftToRight := false
		// fmt.Printf("scanning %#v\n", values[i])

		for j := 9; j >= 0; j-- {
			answer := 0
			index := strings.Index(values[i], strconv.Itoa(j))

			fmt.Printf("found %#v in position %#v\n", j, index)
			if index >= 0 {
				if counter == 0 {
					firstnumber = j
					firstindex = index
					counter += 1
					for k := 9; k >= 0; k-- {
						if firstindex+1 != len(values[i]) {
							secondindex = strings.Index(values[i][1+index:], strconv.Itoa(k))
							if secondindex >= 0 {
								fmt.Printf("in %#v again found %#v in position %#v\n", values[i][1+index:], k, secondindex)
								// counter += 1
								secondnumber = k
								leftToRight = true
								break
							}
						}
					}
					index = -1
				} else if counter == 1 {
					if !leftToRight {
						secondnumber = j
						secondindex = index
					}
					index = -1

					if leftToRight {
						answer = firstnumber*10 + secondnumber
					} else {
						answer = firstnumber + secondnumber*10
					}
					fmt.Printf("answer: %#v\n", answer)
					total += answer
					break
				}
			}
		}
	}
	fmt.Printf("total is %#v", total)
}

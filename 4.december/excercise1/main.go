package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sameNumbers = []string{"00", "11", "22", "33", "44", "55", "66", "77", "88", "99"}

func numberArray(lower, upper int) []int {
	var numbers []int
	for i := lower; i <= upper; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func sizeChecker(pw int) bool {
	if pw < 10000 {
		return false
	}
	if pw > 999999 {
		return false
	}
	return true
}

func adjacentChecker(pw int) bool {
	sint := strconv.Itoa(pw)
	for _, samenumbers := range sameNumbers {
		if strings.Contains(sint, samenumbers) {
			return true
		}
	}
	return false
}

func smallerthanPrevious(pw int) bool {
	sint := strconv.Itoa(pw)
	for i, _ := range sint {
		if (i+1 < len(sint)) && (sint[i] > sint[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	var correctPasswords []int
	passwordRange := numberArray(20000, 200000)
	for _, values := range passwordRange {
		if smallerthanPrevious(values) && adjacentChecker(values) && sizeChecker(values) {
			correctPasswords = append(correctPasswords, values)
		}
	}
	fmt.Printf("Your number of correct passwords is: %d   \n", len(correctPasswords))

}

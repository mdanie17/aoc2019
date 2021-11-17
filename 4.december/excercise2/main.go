package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sameNumbers = []string{"00", "11", "22", "33", "44", "55", "66", "77", "88", "99"}
var toomanySameNumbers = []string{
	"000000", "111111", "222222", "333333", "444444", "555555", "666666", "777777", "888888", "999999",
	"00000", "11111", "22222", "33333", "44444", "55555", "66666", "77777", "88888", "99999",
	"0000", "1111", "2222", "3333", "4444", "5555", "6666", "7777", "8888", "9999",
	"000", "111", "222", "333", "444", "555", "666", "777", "888", "999",
}

func numberArray(lower, upper int) []int {
	var numbers []int
	for i := lower; i <= upper; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func sizeChecker(pw int) bool {
	if pw < 100000 {
		return false
	}
	if pw > 999999 {
		return false
	}
	return true
}

func adjacentChecker(pw int) bool {
	sint := strconv.Itoa(pw)
	for _, samenumbers := range toomanySameNumbers {
		sint = strings.Replace(sint, samenumbers, " ", -1)
	}
	for _, samenumbers := range sameNumbers {
		if strings.Contains(sint, samenumbers) {
			return true
		}
	}
	return false
}

func smallerthanPrevious(pw int) bool {
	sint := strconv.Itoa(pw)
	for i := range sint {
		if (i+1 < len(sint)) && (sint[i] > sint[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	var correctPasswords []int
	passwordRange := numberArray(206938, 679128)
	for _, values := range passwordRange {
		if smallerthanPrevious(values) && adjacentChecker(values) && sizeChecker(values) {
			correctPasswords = append(correctPasswords, values)
		}
	}
	fmt.Printf("Your number of correct passwords is: %d   \n", len(correctPasswords))

}

package main

import (
	"fmt"
	"strconv"
)

type wire struct {
	x    int
	y    int
	dist int
}

func main() {
	testwire1 := []string{"R8", "U5", "L5", "D3"}
	testwire2 := []string{"U7", "R6", "D4", "L4"}
	_ = testwire1
	wire1 := wire{}
	wire2 := wire{0, 0, 0}
	wire1history := []wire{}
	wire2history := []wire{}
	for _, v := range testwire1 {
		l, n := stringSplitter(v)
		wire1.moveChecker(l, n)
		wire1history = append(wire1history, wire1)
	}
	fmt.Println(wire1history)
	//
	for _, v2 := range testwire2 {
		l2, n2 := stringSplitter(v2)
		wire2.moveChecker(l2, n2)
		wire2history = append(wire2history, wire2)
	}
	fmt.Println(wire2history)

	// fmt.Printf("x: %d  y: %d   dist: %d \n", wire1.x, wire1.y, wire1.dist)
	// fmt.Printf("x: %d  y: %d   dist: %d \n", wire2.x, wire2.y, wire2.dist)
}

func stringSplitter(s string) (letters, numbers string) {

	var number, letter []rune
	for _, v := range s {
		switch {
		case v >= 'A' && v <= 'Z':
			letter = append(letter, v)

		case v >= '0' && v <= '9':
			number = append(number, v)
		}

	}
	return string(letter), string(number)
}

func (w *wire) moveChecker(direction string, length string) {
	lengthint, _ := strconv.Atoi(length)

	if direction == "U" {
		w.y += lengthint
	}
	if direction == "D" {
		w.y -= lengthint
	}
	if direction == "L" {
		w.x -= lengthint
	}
	if direction == "R" {
		w.x += lengthint
	}
	w.dist = w.x + w.y
}

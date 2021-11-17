package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type vec struct {
	dir  rune
	dist int
}

type coord struct {
	x, y int
}

func readVecPaths(filePath string) [][]vec {
	txt, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.TrimSpace(string(txt))
	var vecPaths [][]vec
	for _, line := range strings.Split(lines, "\n") {
		var vecPath []vec
		for _, move := range strings.Split(line, ",") {
			var v vec
			_, _ = fmt.Sscanf(move, "%c%d", &v.dir, &v.dist)
			vecPath = append(vecPath, v)
		}
		vecPaths = append(vecPaths, vecPath)
	}
	return vecPaths
}

func moveChecker(vecPath []vec) []coord {
	var path []coord
	var cur coord
	for _, v := range vecPath {
		var dim *int
		d := 0
		switch v.dir {
		case 'U':
			dim = &cur.y
			d = +1
		case 'D':
			dim = &cur.y
			d = -1
		case 'R':
			dim = &cur.x
			d = +1
		case 'L':
			dim = &cur.x
			d = -1
		}
		for n := v.dist; n > 0; n-- {
			*dim += d
			path = append(path, cur)
		}
	}
	return path
}

func findIntersects(wire1cords, wire2cords []coord) []coord {
	coords := make(map[coord]bool)
	for _, c := range wire1cords {
		coords[c] = true
	}
	var intersections []coord
	for _, c := range wire2cords {
		if coords[c] {
			intersections = append(intersections, c)
		}
	}
	return intersections
}

func dist(c coord) int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func closestIntersect(intersects []coord) int {
	closestDist := 99999
	for _, coord := range intersects {
		d := dist(coord)
		if d < closestDist {
			closestDist = d
		}
	}
	return closestDist
}

func intersectIndex(to coord, path []coord) int {
	for i, cur := range path {
		if cur == to {
			return i + 1
		}
	}
	return 0
}

func fastestIntersect(coords []coord, path1, path2 []coord) int {
	moves := 999999
	for _, c := range coords {
		sum := intersectIndex(c, path1) + intersectIndex(c, path2)
		if sum < moves {
			moves = sum
		}
	}
	return moves
}

func main() {
	vecPaths := readVecPaths("inputfull.txt")
	wire1 := moveChecker(vecPaths[0])
	wire2 := moveChecker(vecPaths[1])
	intersects := findIntersects(wire1, wire2)
	fmt.Println(fastestIntersect(intersects, wire1, wire2))
}

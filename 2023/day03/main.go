package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func (this Point) checkPoint(d Point) Point {
	return Point{this.X + d.X, this.Y + d.Y}
}

func part1(file []byte) {
	symbols := map[Point]string{}
	parts := map[Point]string{}
	for y, s := range strings.Fields(string(file)) {
		for x, r := range s {
			if r != 46 && (r < 48 || r > 57) {
				symbols[Point{x, y}] = string(r)
			}
			if r > 47 && r < 58 {
				parts[Point{x, y}] = string(r)
			}
		}
	}
	fmt.Println(symbols)
	fmt.Println(parts)

	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	found := map[Point]string{}
	//TODO: For each symbol coordinate, search around and grab coordinate
	// if that coordinate exists in the parts map
	for k, _ := range symbols {
		//fmt.Printf("Symbol Point: %v:\n", k)
		for _, d := range directions {
			check := Point{k.X, k.Y}.checkPoint(d)

			if s, exists := parts[check]; exists {
				found[check] = s
			}
		}
	}
	fmt.Println(found)
	y := 0
	x := -1
	for k, _ := range found {
		fmt.Println(k)
		//Walk Left

		//for x := -1; test := parts[k.checkPoint(Point{x,y})]; x--
		check := k.checkPoint(Point{x, y})
		fmt.Println(check)
		x--
	}

}

func part2() {

}

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(file)

}

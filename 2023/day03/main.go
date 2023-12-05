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

func part1(file []byte) {
	symbols := map[Point]string{}
	parts := map[Point]string{}
	for y, s := range strings.Fields(string(file)) {
		for x, r := range s {
			if r != 46 && (r < 48 || r > 57) {
				symbols[Point{x, y}] = string(r)
			}
			if  r > 47 && r < 58 {
				parts[Point{x, y}] = string(r)
			}
		}
	}
	fmt.Println(symbols)
	fmt.Println(parts)

	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	//TODO: For each symbol coordinate, search around and grab coordinate
	// if that coordinate exists in the parts map
	for k, _ := range symbols {
		fmt.Println(k)
		for _, d := range directions {
			if _, exists := symbols[d] {
				fmt.Println()
			}
		}
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
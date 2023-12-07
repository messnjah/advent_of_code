package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type FoundPoint struct {
	X, Y Point
}

func (this Point) checkPoint(d Point) Point {
	return Point{this.X + d.X, this.Y + d.Y}
}

func isCharNumber(s rune) bool {
	if s > 47 && s < 58 {
		return true
	}
	return false
}

func part1(file []byte) {
	symbols := map[Point]string{}
	parts := map[Point]string{} 
	for y, s := range strings.Fields(string(file)) {
		for x, r := range s {
			if r != 46 && (!isCharNumber(r)) {
				symbols[Point{x, y}] = string(r)
			}
			if isCharNumber(r) {
				parts[Point{x, y}] = string(r)
			}
		}
	}

	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	found := map[Point]string{}
	for k, _ := range symbols {
		for _, d := range directions {
			check := Point{k.X, k.Y}.checkPoint(d)

			if s, exists := parts[check]; exists {
				found[check] = s
			}
		}
	}

	numbers := map[FoundPoint]int{}
	for k, origin := range found {
		y := 0
		x_left := -1
		x_right := 1
		isNumberLeft := true
		isNumberRight := true
		start := Point{}
		end := Point{}
		for isNumberLeft == true || isNumberRight == true {
			check_left := k.checkPoint(Point{x_left, y})
			check_right := k.checkPoint(Point{x_right, y})
			if s, exists := parts[check_left]; exists && isNumberLeft {
				origin = s + origin
				x_left--
			} else {
				isNumberLeft = false
			}
			if s, exists := parts[check_right]; exists && isNumberRight {
				origin = origin + s
				x_right++
			} else {
				isNumberRight = false
			}
			start = check_left.checkPoint(Point{1,0})
			end = check_right.checkPoint(Point{-1,0})
		}
		found_point := FoundPoint{start,end}
		if _, exists := numbers[found_point]; !exists {
			num, _ := strconv.Atoi(origin)
			numbers[found_point] = int(num)
		}
	}
	total := 0
	for _, i := range numbers {
		total += i 
	}

	fmt.Println(total)
}

func part2() {

}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(file)

}

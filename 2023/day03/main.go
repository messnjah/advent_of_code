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

func getFoundNumber (found map[Point]string, parts map[Point]string) map[FoundPoint]int {
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
	return numbers
}

func getFoundGear (found map[Point][]Point, parts map[Point]string) {
	foundGears := map[Point]map[FoundPoint]int{}
	total := 0
	for k, list := range found {
		foundGears[k] = make(map[FoundPoint]int)
		for _, p := range list {
			y := 0
			x_left := -1
			x_right := 1
			isNumberLeft := true
			isNumberRight := true
			start := Point{}
			end := Point{}
			origin, _ := parts[p]
			for isNumberLeft == true || isNumberRight == true {
				check_left := p.checkPoint(Point{x_left, y})
				check_right := p.checkPoint(Point{x_right, y})
				
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
			if _, exists := foundGears[k][found_point]; !exists {
				num, _ := strconv.Atoi(origin)
				foundGears[k][found_point] = int(num)
			}
		}
		if len(foundGears[k]) > 2 {
			delete(foundGears, k)
		} else {
			keys := [2]FoundPoint{}
			i := 0
			for fk, _ := range foundGears[k] {
				keys[i] = fk
				i++
			}
			ratio := foundGears[k][keys[0]] * foundGears[k][keys[1]]
			total += ratio
		}
	}
	fmt.Println(total)
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
	gear_ratios := map[Point][]Point{} //P2
	for k, symbol := range symbols {
		for _, d := range directions {
			check := Point{k.X, k.Y}.checkPoint(d)
			
			if s, exists := parts[check]; exists {
				found[check] = s
				//P2
				if symbol == "*" { 
					gear_ratios[k] = append(gear_ratios[k], check)
				}
			}
		}
		//P2
		if len(gear_ratios[k]) < 2 { 
			delete(gear_ratios, k)
		}
	}

	numbers := getFoundNumber(found, parts)
	total := 0
	for _, i := range numbers {
		total += i 
	}

	fmt.Println(total)
	//P2
	getFoundGear(gear_ratios, parts)
}

func part2() {
	//Look at getFoundGear function
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(file)

}

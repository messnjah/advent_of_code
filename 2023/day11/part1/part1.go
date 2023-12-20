package main

import (
	"advent_of_code/util"
	"fmt"
	"math"
	"strings"
	"time"
)

type Point struct {
	Y, X int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func addSpace (lines []string) ([]int, []int) {
	//height := len(lines)
	emptyRow := []int{}
	emptyCol := []int{}

	for y, line := range lines {
		if !strings.Contains(line, "#") {
			emptyRow = append(emptyRow, y)
		}
	}

	for x, _ := range lines[0] {
		empty := true
		for y, _ := range lines {
			if lines[y][x] == 35 {
				empty = false
			}
		}
		if empty {
			emptyCol = append(emptyCol, x)
		}
	}

	return emptyRow, emptyCol
}

func processFile(lines []string) {
	//fmt.Println(lines)
	emptyRow, emptyCol := addSpace(lines) //TODO: Dont modify grid, track the indices
	fmt.Println(emptyRow)
	fmt.Println(emptyCol)
	//scale := 2
	
	galaxies := []Point{}
	for y, line := range lines {
		for x, item := range line {
			if item == 35 {
				galaxies = append(galaxies, Point{y,x})
			} 
		}		
	}
	//fmt.Println(lines)
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	fmt.Println(galaxies)

	total := 0
	cnt:=0
	for i, point := range galaxies {
		for _, x := range galaxies[i+1:] {
			minY := math.Min(float64(point.Y),float64(x.Y))
			maxY := math.Max(float64(point.Y),float64(x.Y))
			minX := math.Min(float64(point.X),float64(x.X))
			maxX := math.Max(float64(point.X),float64(x.X))
			scaleY := 0
			scaleX := 0

			for _, row := range emptyRow {
				if int(minY) < row && int(maxY) > row {
					scaleY++
				}
			}
			for _, col := range emptyCol {
				if int(minX) < col && int(maxX) > col {
					scaleX++
				}
			}

			a := Abs(x.X - point.X) + scaleX
			b := Abs(x.Y - point.Y) + scaleY
			//fmt.Printf("%v: Points: %v -> %v, Value: %v, YScale: %v, XScale: %v\n",cnt,point,x,a+b,scaleY,scaleX)
			total += (a + b)
			//fmt.Println(total)
			cnt++
		}
	}
	fmt.Println(total)
}

func main() {
	start := time.Now()
	lines := util.GetFile("../input.txt")
	processFile(lines)
	fmt.Println(time.Since(start))
}
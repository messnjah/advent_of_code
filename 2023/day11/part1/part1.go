package main

import (
	"advent_of_code/util"
	"fmt"
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

func addSpace (lines []string) []string {
	indices := []int{}
	for y, line := range lines {
		if !strings.Contains(line, "#") {
			indices = append(indices, y)
		}
	}
	for idx, i := range indices {
		if idx == 0 {
			lines = append(lines[:i], lines[i:]... )
		} else {
			i++
		}
		lines = append(lines[:i+1], lines[i:]... )
	}

	return lines
}

func processFile(lines []string) {
	//fmt.Println(lines)
	lines = addSpace(lines) //TODO: Need to expand columns... FML
	
	galaxies := []Point{}
	for y, line := range lines {
		for x, item := range line {
			if item == 35 {
				galaxies = append(galaxies, Point{y,x})
			} 
		}		
	}
	//fmt.Println(lines)
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println(galaxies)

	total := 0
	cnt:=0
	for i, point := range galaxies {
		for _, x := range galaxies[i+1:] {
			a := Abs(x.X - point.X)
			b := Abs(x.Y - point.Y)
			fmt.Printf("%v: Points: %v -> %v, Value: %v\n",cnt,point,x,Abs(a)+Abs(b))
			total += (a + b)
			//fmt.Println(total)
			cnt++
		}
	}
	fmt.Println(total)
}

func main() {
	start := time.Now()
	lines := util.GetFile("../test.txt")
	processFile(lines)
	fmt.Println(time.Since(start))
}
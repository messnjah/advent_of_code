package main

import (
	"advent_of_code/util"
	"fmt"
)

var Pipes = map[string][]Point {
	"|": {{-1,0},{1,0}},
	"-": {{0,-1},{0,1}},
	"L": {{-1,0},{0,1}},
	"J": {{0,-1},{1,0}},
	"7": {{0,-1},{-1,0}},
	"F": {{-1,0},{0,1}},
}

type Node struct {
	id Point
	Edges []Point
	Counted bool
}

type Point struct {
	Y,X int
}

type Grid map[Point]Node

func processFile(lines []string) {
	grid := Grid{}
	for y, line := range lines {
		for x, item := range line {
			if item != 46 {
				node := Node{}
				node.id = Point{y,x}
				node.Edges = Pipes[string(item)]
				grid[node.id] = node
			}
		}
	}
	fmt.Println(grid)
}

func main() {
	lines := util.GetFile("../test.txt")
	processFile(lines)
}
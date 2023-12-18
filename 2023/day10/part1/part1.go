package main

import (
	"advent_of_code/util"
	"fmt"
)

var Pipes = map[string][]Point {
	"|": {{-1,0},{1,0}},
	"-": {{0,-1},{0,1}},
	"L": {{-1,0},{0,1}},
	"J": {{0,-1},{-1,0}},
	"7": {{0,-1},{1,0}},
	"F": {{1,0},{0,1}},
}

type Node struct {
	id Point
	Edges []Point
	Counted bool
	Steps int
}

type Point struct {
	Y,X int
}

type Grid map[Point]*Node

func (this Point) checkPoint(d Point) Point {
	return Point{this.Y + d.Y, this.X + d.X}
}

func (this *Node) String() string { // Stringer interface to print values at memory address
	return fmt.Sprintf("Edges: %d, Steps: %d\n", this.Edges, this.Steps)
}

func startPipe (grid Grid, start Point) []Point {
	locations := []Point{{-1,0},{0,-1},{1,0},{0,1}}
	foundEdges := []Point{}
	for _, point := range locations {
		i := start.checkPoint(point)
		if node, exists := grid[i]; exists {
			for _, edge := range node.Edges {
				if node.id.checkPoint(edge) == start {
					foundEdges = append(foundEdges, point)
				}
			}
		}
	}
	return foundEdges
}

func (this *Node) stepsFromStart (grid Grid, start, previous Point, steps int) {
	/* if this.id == start {
		this.Steps = steps
		this.Counted = true
	} else {
		steps++
		this.Steps = steps
		this.Counted = true
	} */

	steps++
	this.Steps = steps
	this.Counted = true
	for _, edge := range this.Edges {

		next := this.id.checkPoint(edge)
		if next != start && next != previous {
			fmt.Printf("Previous: %v, Current: %v, Next: %v\n", previous, this.id, next)
			grid[next].stepsFromStart(grid, start, this.id, steps)
		} else if next == start && next != previous {
			grid[start].Steps = steps
		}
	}
}

func processFile(lines []string) {
	grid := Grid{}
	start := Point{}
	for y, line := range lines {
		for x, item := range line {
			if item != 46 {
				node := &Node{}
				if item == 83 {
					start = Point{y,x}
					node.id = start
					grid[start] = node
					continue
				}
				node.id = Point{y,x}
				node.Edges = Pipes[string(item)]
				grid[node.id] = node
			}
		}
	}
	grid[start].Edges = startPipe(grid, start)
	fmt.Println(grid)
	grid[start].stepsFromStart(grid, start, start, 0)
	
	fmt.Println(grid[start].Steps / 2)
}

func main() {
	lines := util.GetFile("../input.txt")
	processFile(lines)
}
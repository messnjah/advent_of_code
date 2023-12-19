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
	Path []Point
}

type Point struct {
	Y,X int
}

type Grid map[Point]*Node

func (this Point) checkPoint(d Point) Point {
	return Point{this.Y + d.Y, this.X + d.X}
}

func (this *Node) String() string { // Stringer interface to print values at memory address
	return fmt.Sprintf("Edges: %d, Counted: %v\n", this.Edges, this.Counted)
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
	
	steps++
	this.Steps = steps
	this.Counted = true
	for _, edge := range this.Edges {

		next := this.id.checkPoint(edge)
		if next != start && next != previous {
			//fmt.Printf("Previous: %v, Current: %v, Next: %v\n", previous, this.id, next)
			grid[start].Path = append(grid[start].Path, next)
			grid[next].stepsFromStart(grid, start, this.id, steps)
		} else if next == start && next != previous {
			grid[start].Steps = steps
		}
		if this.id == start {
			break
		}
	}
}

func shoelace(grid Grid, path []Point) { // https://en.wikipedia.org/wiki/Shoelace_formula
	// vertices := []Point{}
	// for point, node := range grid { // Get all points that from the pipe
	// 	if node.Counted == true {
	// 		vertices = append(vertices, point)
	// 	}
	// }

	// Find Center
	// center := Point{}
	// xsum := 0
	// ysum := 0
	// for _, point := range vertices {
	// 	xsum += point.X
	// 	ysum += point.Y
	// }
	// center.X = xsum / len(vertices)
	// center.Y = ysum / len(vertices)

	// Find each points distance from the center
	// angles := []float64{}
	// indices := make([]int, len(vertices)) // Create index copy to sort on
	// for idx, point := range vertices {
	// 	indices[idx] = idx
	// 	angle := math.Atan2(float64(point.Y - center.Y),float64(point.X - center.X))
	// 	angles = append(angles, angle)
	// }
	
	// Sort the index copy based on the angles values
	// sort.SliceStable(indices, func(i, j int) bool {
	// 	return angles[indices[i]] > angles[indices[j]]
	// })
	// Create new slice with sorted points
	// result := make([]Point, len(vertices))
	// for idx := range result {
	// 	result[idx] = vertices[indices[idx]]
	// }
	//fmt.Println(result)

	sum1 := 0
	for i:=0; i < len(path); i++ {
		if i == len(path) - 1 {
			sum1 += (path[i].Y + path[0].Y) * (path[i].X - path[0].X)
			//sum1 += (result[i].X * result[0].Y) - (result[i].Y * result[0].X)
			//sum2 += result[i].Y * result[0].Y
		} else {
			sum1 += (path[i].Y + path[i+1].Y) * (path[i].X - path[i+1].X)
			//sum2 += result[i].Y * result[i+1].Y
		}
	}
	area := Abs(sum1)/ 2

	fmt.Println(area)

	// https://en.wikipedia.org/wiki/Pick's_theorem
	inner := area - (len(path) / 2) + 1
	fmt.Println(inner)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func processFile(lines []string) {
	grid := Grid{}
	//nonPipe := map[Point]
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
			} else {

			}
		}
	}
	grid[start].Edges = startPipe(grid, start)
	grid[start].Path = append(grid[start].Path, start)
	
	grid[start].stepsFromStart(grid, start, start, 0)
	//fmt.Println(grid[start].Path)
	fmt.Println(grid[start].Steps / 2)

	shoelace(grid, grid[start].Path)
}

func main() {
	lines := util.GetFile("../input.txt")
	processFile(lines)
}
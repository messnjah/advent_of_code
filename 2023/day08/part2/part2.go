package main

import (
	"fmt"
	"strings"

	"advent_of_code/util"
)

type Node struct {
	left string
	right string
}

type Tree map[string]Node


func processFile(lines []string) {
	direct := strings.Split(lines[0],"")
	tree := Tree{}
	for _, line := range lines[2:] {
		node := Node{}
		split := strings.Split(line, "=")
		id := strings.TrimSpace(split[0])
		routes := strings.Split(strings.Trim(split[1], " ()"),",")
		node.left = strings.TrimSpace(routes[0])
		node.right = strings.TrimSpace(routes[1])
		tree[id] = node
	}
	
	steps := 0
	next := "AAA"
	for x := 0; next != "ZZZ"; x++ {
		if x == len(direct) {
			x = 0
		}
		
		if direct[x] == "R" {
			next = tree[next].right
		} else if direct[x] == "L" {
			next = tree[next].left
		}
		steps++
	}
	fmt.Println(steps)
}

func main() {
	lines := util.GetFile("input.txt")
	processFile(lines)
}
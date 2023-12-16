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

func locations(id string, compare byte) bool {
	if id[2] == compare {
		return true
	}
	return false
}

func processFile(lines []string) {
	direct := strings.Split(lines[0],"")
	tree := Tree{}
	starts := []string{}
	for _, line := range lines[2:] {
		node := Node{}
		split := strings.Split(line, "=")
		id := strings.TrimSpace(split[0])
		routes := strings.Split(strings.Trim(split[1], " ()"),",")
		node.left = strings.TrimSpace(routes[0])
		node.right = strings.TrimSpace(routes[1])
		tree[id] = node
		if locations(id, 65) {
			starts = append(starts, id)
		}
	}
	fmt.Println(starts)
	endSteps := []int{}
	for _, start := range starts {
		steps := 0
		for x := 0; !locations(start,90); x++ {
			if x == len(direct) {
				x = 0
			}
			
			if direct[x] == "R" {
				start = tree[start].right
			} else if direct[x] == "L" {
				start = tree[start].left
			}
			steps++
		}
		endSteps = append(endSteps, steps)
	}
	fmt.Println(endSteps)
	finalSteps := util.LCM(endSteps[0], endSteps[1:])
	fmt.Println(finalSteps)
}

func main() {
	lines := util.GetFile("../input.txt")
	processFile(lines)
}
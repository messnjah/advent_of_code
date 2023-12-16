package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent_of_code/util"
)

func findDiff(line []int) []int {
	output := []int{}
	for i:=0;i < len(line)-1; i++ {
		first := line[i+1]
		next := line[i]
		diff := first - next
		output = append(output, diff)
	}
	return output
}

func allZeros(line []int) bool {
	for _, i := range line {
		if i != 0 {
			return false
		}
	}
	return true
}

func processFile(lines []string) {
	total := 0
	for _, line := range lines {
		outerList := [][]int{}
		listInt := []int{}
		items := strings.Fields(line)
		for i:=0;i < len(items); i++ {
			item, _ := strconv.Atoi(items[i])
			listInt = append(listInt, item)
		}
		outerList = append(outerList, listInt)
		set := findDiff(listInt)
		outerList = append(outerList, set)
		for !allZeros(set) {
			set = findDiff(set)
			outerList = append(outerList, set)
		}

		nextNum := 0
		slices.Reverse(outerList)
		cnt := 0
		for _, x := range outerList {
			if allZeros(x) {
				continue
			}
			nextNum = outerList[cnt+1][0] - nextNum
			cnt++
		}
		total += nextNum

		
	
	}
	fmt.Println(total)
	
}

func main() {
	lines := util.GetFile("../input.txt")
	processFile(lines)
}
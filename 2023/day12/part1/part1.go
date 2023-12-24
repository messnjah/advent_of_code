package main

import (
	"advent_of_code/util"
	"fmt"
	"strings"
	"time"
)

func processFile(lines []string) {
	
	for _, line := range lines {
		fields := strings.Fields(line)
		groups := strings.Split(fields[1],",")

		
	}
}

func main() {
	start := time.Now()
	lines := util.GetFile("../test.txt")
	processFile(lines)
	fmt.Println(time.Since(start))
}
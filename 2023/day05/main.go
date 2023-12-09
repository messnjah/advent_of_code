package main

import (
	"fmt"
	"os"
	"strings"
)

type Detail struct {
	src int
	dst int
	length int
}

type Row struct {
	details []Detail
}

type Maps map[string]Row

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(string(file), "\n")
}

func isCharNumber(s byte) bool {
	if s > 47 && s < 58 {
		return true
	}
	return false
}

func processFile(lines []string) {
	alm := Maps{}
	for line, s := range lines {
		if line == 0 {
			
		}
		if s[0] == 13 {
			continue
		} else {}
	}
	fmt.Println(alm)
}

func main() {
	file := "test.txt"
	lines := readLines(file)
	processFile(lines)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	return strings.Split(string(file), "\n")
}

func processFile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {

	}
	test := strings.Split(string(file), "\r\n")
	//test2 := strings.Split(string(file), ":")
	for y, s := range test {
		if strings.Contains(s, "\r") {
			fmt.Println(s)
		} else {
			fmt.Printf("%v %v\n", y, s)
		}

	}
	//fmt.Println(test2[1])
}

func main() {
	file := "test.txt"
	lines := readLines(file)
	fmt.Println(lines[0])

	processFile(file)
}

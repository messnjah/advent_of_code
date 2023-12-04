package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1(s string) {
	colon_idx := strings.Index(s, ":")
	//game_num := s[5:colon_idx]
	//fmt.Println(game_num)

	restofstring := s[colon_idx+2 : len(s)]
	set_num := strings.Count(restofstring, ";") + 1
	fmt.Println(set_num)

}

func part2(s string) {
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_scanner := bufio.NewScanner(file)

	for file_scanner.Scan() {
		//fmt.Println(file_scanner.Text())
		part1(file_scanner.Text())
	}

}

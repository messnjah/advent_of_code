package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(s string) int {
	max_red := 12
	max_green := 13
	max_blue := 14
	colon_idx := strings.Index(s, ":")
	game_num := s[5:colon_idx]
	//fmt.Println(game_num)

	red := 0
	green := 0
	blue := 0
	//words := []string{"red","green","blue"}

	restofstring := s[colon_idx+2 : len(s)]
	for _, set := range strings.Split(restofstring, ";") {
		for _, draw := range strings.Split(set, ",") {
			results := strings.Split(strings.TrimSpace(draw), " ")
			color := results[1]
			color_num, _ := strconv.Atoi(results[0])
			//fmt.Println("Game ID: " + game_num + " Number: " + results[0] + " Color: " + results[1])
			if color == "red" {
				if color_num > red {
					red = color_num
				}
			}
			if color == "green" {
				if color_num > green {
					green = color_num
				}
			}
			if color == "blue" {
				if color_num > blue {
					blue = color_num
				}
			}
		}
	}
	if red > max_red || green > max_green || blue > max_blue {
		//fmt.Println("Game: " + game_num + " Not Possible")
		return 0
	} else {
		//fmt.Println("Game: " + game_num + " Possible")
		game_num, _ := strconv.Atoi(game_num)
		return game_num
	}

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
	part1_total := 0

	for file_scanner.Scan() {
		//fmt.Println(file_scanner.Text())
		part1_total += part1(file_scanner.Text())

	}
	fmt.Println(part1_total)
}

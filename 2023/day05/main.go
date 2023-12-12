package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Detail struct {
	src    int
	dst    int
	length int
	srcEnd int
	dstEnd int
}

type Maps map[string]*[]Detail

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n\n") //Windows :(
}

func isCharNumber(s byte) bool {
	if s > 47 && s < 58 {
		return true
	}
	return false
}

func processFile(lines []string) (Maps, []int) {
	alm := Maps{}
	seeds := []int{}
	for line, s := range lines {
		if line == 0 {
			for _, num := range strings.Fields(strings.Split(s, ":")[1]) {
				i, _ := strconv.Atoi(num)
				seeds = append(seeds, i)
			}
			continue
		}
		splitMap := strings.Split(s, "map:")
		mapName := strings.TrimSpace(splitMap[0])
		mapNumbers := strings.Split(strings.TrimSpace(splitMap[1]), "\n")
		rows := []Detail{}
		for _, row := range mapNumbers {
			detail := Detail{}
			numbers := strings.Fields(row)
			detail.dst, _ = strconv.Atoi(numbers[0])
			detail.src, _ = strconv.Atoi(numbers[1])
			detail.length, _ = strconv.Atoi(numbers[2])
			detail.srcEnd = detail.src + detail.length - 1
			detail.dstEnd = detail.dst + detail.length - 1
			rows = append(rows, detail)
		}
		alm[mapName] = &rows
	}

	return alm, seeds
}

func part1(alm Maps, seeds []int) {
	soilMap := alm["fertilizer-to-water"]
	seedStart := 0
	seedEnd := 0
	soilStart := 0
	soilEnd := 0

	
	for idx, detail := range *soilMap {
		if idx == 0 {
			seedStart = detail.src
			seedEnd = detail.src + detail.length - 1
			soilStart = detail.dst
			soilEnd = detail.dst + detail.length - 1
		} else {
			if detail.src < seedStart {
				seedStart = detail.src
			}
			if detail.src + detail.length - 1 > seedEnd {
				seedEnd = detail.src + detail.length - 1
			}
			if detail.dst < soilStart {
				soilStart = detail.dst
			}
			if detail.dst + detail.length - 1 > soilEnd {
				soilEnd = detail.dst + detail.length - 1 
			}
		}
	}

	fmt.Println(soilMap)
	fmt.Printf("seed range: %v - %v\nsoil range: %v - %v\n",seedStart,seedEnd,soilStart,soilEnd)
}

func main() {
	file := "test.txt"
	lines := readLines(file)
	alm, seeds := processFile(lines)
	part1(alm, seeds)

}

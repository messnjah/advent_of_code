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
			rows = append(rows, detail)
		}
		alm[mapName] = &rows
	}

	return alm, seeds
}

func part1(alm Maps, seeds []int) {
	//seedStart := 0
	soilMap := alm["seed-to-soil"]
	for _, detail := range *soilMap {
		fmt.Println(detail)
	}

	fmt.Println(seeds)
	for key := range alm {
		fmt.Println(key)
	}
}

func main() {
	file := "test.txt"
	lines := readLines(file)
	alm, seeds := processFile(lines)
	part1(alm, seeds)

}

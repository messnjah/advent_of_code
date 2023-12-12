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
	diff int
}

type Maps map[string]*[]Detail

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n\n") //Windows :(
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
			detail.diff = detail.dst - detail.src
			rows = append(rows, detail)
		}
		alm[mapName] = &rows
	}

	return alm, seeds
}

func part1(alm Maps, seeds []int) {
	mapsList := [7]string {"seed-to-soil","soil-to-fertilizer","fertilizer-to-water","water-to-light","light-to-temperature","temperature-to-humidity","humidity-to-location" }
	location := 0
	for _, seed := range seeds {
		input := seed
		output := 0
		found := false
		

		for idx, name := range mapsList {
			subMap := alm[mapsList[idx]]

			for _, item := range *subMap {
				if input >= item.src && input <= item.srcEnd {
					output = input + item.diff
					//fmt.Printf("input: %v, %v: %v\n",input, name, output)
					found = true
				} 
			}
	
			if !found {
				output = input
				//fmt.Printf("input: %v, %v: %v\n",input, name, output)
			}
			found = false

			if name == "humidity-to-location" && (location == 0 || output < location) {
				location = output
			}
			input = output	
		}	
	}
	fmt.Println(location)
}

func part2(alm Maps, seeds []int) {
	mapsList := [7]string {"seed-to-soil","soil-to-fertilizer","fertilizer-to-water","water-to-light","light-to-temperature","temperature-to-humidity","humidity-to-location" }
	location := 0
	newSeeds := [][]int{}
	for i:=0; i < len(seeds); i += 2 {
		temp := []int{}
		result := seeds[i:i+2]
		temp = append(temp, result...)
		newSeeds = append(newSeeds, temp)
	}
	fmt.Println(newSeeds)

	for _, seedRange := range newSeeds {
		for i := seedRange[0]; i < seedRange[0] + seedRange[1]; i++ {
			input := i
			output := 0
			found := false			

			for _, name := range mapsList {
				subMap := alm[name]

				for _, item := range *subMap {
					if input >= item.src && input <= item.srcEnd {
						output = input + item.diff
						//fmt.Printf("input: %v, %v: %v\n",input, name, output)
						found = true
					} 
				}
		
				if !found {
					output = input
					//fmt.Printf("input: %v, %v: %v\n",input, name, output)
				}
				found = false

				if name == "humidity-to-location" && (location == 0 || output < location) {
					location = output
				}
				input = output	
			}
		}
	}
	fmt.Println(location)
}

func main() {
	file := "test.txt"
	lines := readLines(file)
	alm, seeds := processFile(lines)
	part1(alm, seeds)
	part2(alm, seeds)
}

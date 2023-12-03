package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(s string) (int, error) {
	number_slice := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	idx_locations := [][]string{}

	for _, number := range number_slice {
		if found_number := strings.Contains(s, number); found_number == true {
			results := []string{}
			results = append(results, strconv.Itoa(strings.Index(s, number)), number)
			idx_locations = append(idx_locations, [][]string{results}...)
			last := []string{}
			last = append(last, strconv.Itoa(strings.LastIndex(s, number)), number)
			idx_locations = append(idx_locations, [][]string{last}...)
		}

	}
	sort.Slice(idx_locations, func(i, j int) bool {
		a, _ := strconv.Atoi(idx_locations[i][0])
		b, _ := strconv.Atoi(idx_locations[j][0])
		return a < b
	})
	return strconv.Atoi(idx_locations[0][1] + idx_locations[len(idx_locations)-1][1])
}

func part2(s string) (int, error) {
	number_map := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9"}
	idx_locations := [][]string{}

	for word, number := range number_map {
		if found_word := strings.Contains(s, word); found_word == true {
			results := []string{strconv.Itoa(strings.Index(s, word)), number}
			idx_locations = append(idx_locations, [][]string{results}...)
			last := []string{}
			last = append(last, strconv.Itoa(strings.LastIndex(s, word)), number)
			idx_locations = append(idx_locations, [][]string{last}...)
		}
		if found_number := strings.Contains(s, number); found_number == true {
			results := []string{}
			results = append(results, strconv.Itoa(strings.Index(s, number)), number)
			idx_locations = append(idx_locations, [][]string{results}...)
			last := []string{}
			last = append(last, strconv.Itoa(strings.LastIndex(s, number)), number)
			idx_locations = append(idx_locations, [][]string{last}...)
		}

	}
	sort.Slice(idx_locations, func(i, j int) bool {
		a, _ := strconv.Atoi(idx_locations[i][0])
		b, _ := strconv.Atoi(idx_locations[j][0])
		return a < b
	})
	return strconv.Atoi(idx_locations[0][1] + idx_locations[len(idx_locations)-1][1])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_scanner := bufio.NewScanner(file)
	part1_total := 0
	part2_total := 0
	for file_scanner.Scan() {
		part1_num, _ := part1(file_scanner.Text())
		part1_total += part1_num
		part2_num, _ := part2(file_scanner.Text())
		part2_total += part2_num
	}
	fmt.Printf("Part 1 total: "+"%d\n", part1_total)
	fmt.Printf("Part 2 total: "+"%d\n", part2_total)
}

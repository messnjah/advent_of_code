package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func reverse_string(s string) string {
	chars := []rune(s)
	for idx, jump := 0, len(chars)-1; idx < jump; idx, jump = idx+1, jump-1 {
		chars[idx], chars[jump] = chars[jump], chars[idx]
	}
	return string(chars)
}

func find_int(s string) (int, string, error) {
	//num_idx :=
	// Iterate over each character in the line
	for idx, char := range s {
		// Check if the character can be converted to an integer
		if _, err := strconv.Atoi(string(char)); err == nil {
			return idx, string(char), nil
		}
	}
	return 0, "", errors.New("No int found")
}

func find_int_name(s string) {
	valid_num := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, number := range valid_num {
		if result := strings.Contains(s, number); result == true {
			fmt.Println(strings.Index(s, number), strings.Index(s, number)+len(number)-1, number)
			// TODO: Find the index locations of the int and word int. Determine which shows up first and grab that number
			// TODO: Create forward array for first number and backward array for last number
		}

	}

}

func part2(s string) (int, error) {
	number_map := map[string]string{
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
			results := []string{strconv.Itoa(strings.Index(s, word)), number} // Get index and number for the word
			idx_locations = append(idx_locations, [][]string{results}...)
		}
		if found_number := strings.Contains(s, number); found_number == true {
			results := []string{} // Get index and number for the word
			results = append(results, strconv.Itoa(strings.Index(s, number)), number)
			idx_locations = append(idx_locations, [][]string{results}...)
		}

	}
	sort.Slice(idx_locations, func(i, j int) bool {
		a, _ := strconv.Atoi(idx_locations[i][0])
		b, _ := strconv.Atoi(idx_locations[j][0])
		return a < b
	})
	fmt.Println(idx_locations[0][1] + idx_locations[len(idx_locations)-1][1])
	fmt.Println(idx_locations)
	return strconv.Atoi(idx_locations[0][1] + idx_locations[len(idx_locations)-1][1])
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_scanner := bufio.NewScanner(file)
	// file_scanner.Split(bufio.ScanLines)
	part1_total := 0
	part2_total := 0
	// Scan the file line by line
	for file_scanner.Scan() {
		_, first_num, _ := find_int(file_scanner.Text())
		_, second_num, _ := find_int(reverse_string(file_scanner.Text()))
		result, _ := strconv.Atoi(first_num + second_num)
		part1_total = part1_total + result
		//fmt.Println(result)
		//find_int_name(file_scanner.Text())
		part2_num, _ := part2(file_scanner.Text())
		part2_total = part2_total + part2_num
	}
	fmt.Println(part1_total)
	fmt.Println(part2_total)

}

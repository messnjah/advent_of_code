package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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

func find_int_name(s string){
	valid_num := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, number := range valid_num {
		if result := strings.Contains(s, number); result == true {
			fmt.Println(strings.Index(s,number), strings.Index(s,number)+len(number)-1, number)
			// TODO: Find the index locations of the int and word int. Determine which shows up first and grab that number
			// TODO: Create forward array for first number and backward array for last number
		}
		
	}

}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_scanner := bufio.NewScanner(file)
	// file_scanner.Split(bufio.ScanLines)
	total := 0
	// Scan the file line by line
	for file_scanner.Scan() {
		first_num_idx, first_num, _ := find_int(file_scanner.Text())
		first_word_idx, first_word := find_int_name(file_scanner.Text()) // TODO: Change word to int
		_, second_num, _ := find_int(reverse_string(file_scanner.Text()))
		result, _ := strconv.Atoi(first_num + second_num)
		total = total + result
		fmt.Println(result)
		//find_int_name(file_scanner.Text())
	}
	fmt.Println(total)
	
}
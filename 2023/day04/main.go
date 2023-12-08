package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type GameCard struct {
	matches int
	explodes []int
	duplicates int
}

func part1(s string) int {
	card := strings.Split(strings.Split(s,":")[1],"|")
	winners := strings.Fields(card[0])
	picked := strings.Fields(card[1])
	found_count := 0
	for _, num := range winners {
		if check := slices.Contains(picked, num); check {
			found_count++
		}
		
	}
	total := 0
	for x := 0; x < found_count; x++ {
		if x == 0 {
			total += 1
		} else {
			total = total * 2
		} 
	}
	return total
}

func part2(scanner *bufio.Scanner) {
	cards := map[int]*GameCard{}
	idx := 1
	for scanner.Scan() {
		cardOutput := &GameCard{}
		card := strings.Split(strings.Split(scanner.Text(),":")[1],"|")
		winners := strings.Fields(card[0])
		picked := strings.Fields(card[1])
		found_count := 0
		for _, num := range winners {
			if check := slices.Contains(picked, num); check {
				found_count++
			}
			
		}
		for x := 1; x < found_count+1; x++ {
			count := idx + x
			cardOutput.explodes = append(cardOutput.explodes, count)
		}
		cardOutput.matches = found_count
		cardOutput.duplicates = 1
		cards[idx] = cardOutput
		idx++
	}

	for x := 1; x < len(cards)+1; x++ {
		if len(cards[x].explodes) > 0 {
			for _, i := range cards[x].explodes {
				cards[i].duplicates += 1 * cards[x].duplicates
			}
		}
	}
	total := 0
	for x := range cards{
		total += cards[x].duplicates
	}
	fmt.Println(total)
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
		part1_total += part1(file_scanner.Text())
	}
	
	fmt.Println(part1_total)
	part2(file_scanner)
	// Need to rework the file scanner, have to comment our part1 for part2 to read
}
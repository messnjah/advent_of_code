package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"advent_of_code/util"
)

var letterMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
}

type Hand struct {
	cards  string
	bet    int
	order  []int
	strength int
}

func handStrength(hand Hand) int {
	cardCounts := []int{}
	for _, card := range hand.cards {
		char := strings.Count(hand.cards, string(card))
		cardCounts = append(cardCounts, char)
	}

	if slices.Contains(cardCounts,5) {
		return 6
	} else if slices.Contains(cardCounts,4) {
		return 5
	} else if slices.Contains(cardCounts,3) {
		if slices.Contains(cardCounts, 2) {
			return 4
		} else {
		return 3
		}
	} else if slices.Contains(cardCounts,2) {
		twos := 0
		for _, item := range cardCounts {
			if item == 2 {
				twos++
			}
		}
		if twos == 4 {
			return 2
		} else {
		return 1
		}
	} else {
	return 0
	}
}

func cardOrder(hand Hand) []int {
	output := []int{}
	for _, char := range hand.cards {
		if value, exists := letterMap[string(char)]; exists {
			output = append(output, value)
		} else {
			num, _ := strconv.Atoi(string(char))
			output = append(output, num)
		}
	}
	return output
}

func applyWilds(hand Hand) int {
	wild := 0
		for _, item := range hand.order {
			if item == 1 {
				wild++
			}
		}
	//fmt.Printf("Wilds: %v, Cards: %v\n", wild, hand.order)
	if wild == 0 {
		return hand.strength
	}

	cardCounts := []int{}
	for _, card := range hand.cards {
		char := strings.Count(hand.cards, string(card))
		cardCounts = append(cardCounts, char)
	}
	//fmt.Printf("Counts: %v, Cards: %v\n", cardCounts, hand.order)
	switch wild {
	case 4:
		return 6
	case 3:
		if slices.Contains(cardCounts,2) { // 5 kind
			return 6
		} else { // 4 kind
			return 5 
		}
	case 2:
		if slices.Contains(cardCounts,3) { // 5 kind
			return 6
		} else if slices.Contains(cardCounts,2) {
			twos := 0
			for _, item := range cardCounts {
				if item == 2 { 
					twos++
				}
			}
			if twos == 4 { // 4 kind
				return 5
			} else { // 3 kind
				return 3
			}
		}
	case 1:
		if slices.Contains(cardCounts,4) { // 5 kind
			return 6
		} else if slices.Contains(cardCounts,3) { // 4 kind
			return 5
		} else if slices.Contains(cardCounts,2) {
			twos := 0
			for _, item := range cardCounts {
				if item == 2 { 
					twos++
				}
			}
			if twos == 4 { // Full House
				return 4
			} else { // 3 kind
				return 3
			}
		} else if slices.Contains(cardCounts,1) { // pair
			return 1
		}
	}
	return hand.strength
}

func processFile(lines []string) {
	hands := []Hand{}

	for _, line := range lines {
		hand := Hand{}
		fields := strings.Fields(line)
		hand.cards  = fields[0] 
		hand.bet, _ = strconv.Atoi(fields[1])
		hand.order = cardOrder(hand)
		hand.strength = handStrength(hand)
		hand.strength = applyWilds(hand)
		hands = append(hands, hand)
	}

	total := 0
	count := 0
	for rank := 0; rank < 7; rank++ {
		group := []Hand{}
		for _, hand := range hands {
			if rank == hand.strength {
				group = append(group, hand)
			}
		}
		sort.SliceStable(group, func(i, j int) bool {
			switch slices.Compare(group[i].order,group[j].order) {
			case -1:
				return true
			case 1:
				return false
			default:
				return false
			}
		})
		//fmt.Println(group)
		for _, item := range group {
			count++
			total += item.bet * count
			if slices.Contains(item.order,1) {
				fmt.Printf("Rank: %v Cards: %v Total: %v\n",count,item,total)
			}
			
		}
	} 
	fmt.Println(count)
	fmt.Println(total)
}

func main() {
	lines := util.GetFile("input.txt")
	processFile(lines)
}

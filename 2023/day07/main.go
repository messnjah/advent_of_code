package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"advent_of_code/util"
)

type Hand struct {
	cards  string
	bet    int
	order  []int
	counts []int
	rank   int
}

func processFile(lines []string) []Hand {
	hands := []Hand{}
	for _, line := range lines {
		hand := Hand{}
		items := strings.Fields(line)
		hand.cards = items[0]
		hand.bet, _ = strconv.Atoi(items[1])

		for _, card := range strings.Split(hand.cards, "") {
			cardNum := 0
			switch card {
			case "A":
				cardNum = 14
			case "K":
				cardNum = 13
			case "Q":
				cardNum = 12
			case "J":
				cardNum = 11
			case "T":
				cardNum = 10
			default:
				cardNum, _ = strconv.Atoi(card)
			}
			hand.order = append(hand.order, cardNum)
			count := strings.Count(hand.cards, card)
			hand.counts = append(hand.counts, count)
		}
		sort.SliceStable(hand.counts, func(i, j int) bool {
			return hand.counts[i] > hand.counts[j]
		})
		hand.rank = getHandRank(hand)
		hands = append(hands, hand)
	}
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].rank > hands[j].rank
	})
	return hands
}

func getHandRank(hand Hand) int {
	rank := 0
	switch hand.counts[0] {
	case 5:
		rank = 7
	case 4:
		rank = 6
	case 3:
		if hand.counts[3] == 2 {
			rank = 5
		} else {
			rank = 4
		}
	case 2:
		if hand.counts[2] == 2 {
			rank = 3
		} else {
			rank = 2
		}
	default:
		rank = 1
	}
	return rank
}

func part1(hands []Hand) {
	length := len(hands)
	fmt.Println(length)
	output := 0
	count := 1
	for i := 1; i <= 7; i++ {
		subSlice := []Hand{}
		for _, hand := range hands {
			if hand.rank != i {
				continue
			}
			subSlice = append(subSlice, hand)
		}
		sort.SliceStable(subSlice, func(i, j int) bool {
			if subSlice[i].order[0] != subSlice[j].order[0] {
				return subSlice[i].order[0] < subSlice[j].order[0]
			} else {
				for x := 1; subSlice[i].order[x] == subSlice[j].order[x]; x++ {
					if subSlice[i].order[x] != subSlice[j].order[x] {
						return subSlice[i].order[x] < subSlice[j].order[x]
					}
				}
				return subSlice[i].order[1] < subSlice[j].order[1]
			}

		})

		fmt.Println(subSlice)
		for _, card := range subSlice {
			fmt.Printf("%v: %v\n", count, card.bet)
			output += card.bet * count
			count++
		}
	}
	fmt.Println(output)
}

func main() {
	lines := util.GetFile("test.txt")
	hands := processFile(lines)
	part1(hands) // Check highest number only cases
}

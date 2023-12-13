package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n") //Windows :(
}

func processFile(lines []string) []Race {
	races := []Race{}
	raceTime := strings.Fields(lines[0])
	raceDist := strings.Fields(lines[1])
	for i := 1; i < len(raceTime); i++ {
		race := Race{}
		race.time, _ = strconv.Atoi(raceTime[i])
		race.distance, _ = strconv.Atoi(raceDist[i])
		races = append(races, race)
	}
	return races
}

func part1(races []Race) {
	results := []int{}
	for _, race := range races {
		wins := 0
		for i := 1; i < race.time; i++ {
			distance := i * (race.time - i)
			if distance > race.distance {
				wins++
			}
		}
		results = append(results, wins)
	}

	output := results[0]
	for i := 1; i < len(results); i++ {
		output *= results[i]
	}

	fmt.Println(output)
}

func part2(races []Race) {
	timeTotal := ""
	distanceTotal := ""
	for _, race := range races {
		outTime := strconv.Itoa(race.time)
		outDist := strconv.Itoa(race.distance)
		timeTotal += outTime
		distanceTotal += outDist
	}
	fmt.Println(timeTotal + " " + distanceTotal)
	timeInt, _ := strconv.Atoi(timeTotal)
	distInt, _ := strconv.Atoi(distanceTotal)
	wins := 0
	for i := 1; i < timeInt; i++ {
		distance := i * (timeInt - i)
		if distance > distInt {
			wins++
		}
	}
	fmt.Println(wins)
}

func main() {
	file := "input.txt"
	lines := readLines(file)
	races := processFile(lines)
	part1(races)
	part2(races)
}

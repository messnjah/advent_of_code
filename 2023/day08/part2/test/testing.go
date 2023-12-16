package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"advent_of_code/util"
)

type mapa struct {
	instructionsLoop string
	positions        map[string][2]string
}

func parseInput(s []string) (m mapa, err error) {
	if len(s) < 3 {
		return m, errors.New("invalid input")
	}
	if strings.Count(s[0], "R")+strings.Count(s[0], "L") != len(s[0]) {
		return m, errors.New("invalid input")
	}
	m.instructionsLoop = s[0]
	m.positions = make(map[string][2]string)
	var p, l, r string
	for idx, v := range s {
		if idx > 1 {
			vs := strings.Split(v, " = ")
			if len(vs) != 2 {
				return m, errors.New("invalid input")
			}
			p = vs[0]
			vss := strings.Split(vs[1], ", ")
			if len(vss) != 2 {
				return m, errors.New("invalid input")
			}
			if vss[0][0] != '(' {
				return m, errors.New("invalid input")
			}
			l = vss[0][1:]
			if vss[1][len(vss[1])-1] != ')' {
				return m, errors.New("invalid input")
			}
			r = vss[1][:len(vss[1])-1]
			m.positions[p] = [2]string{l, r}
		}
	}
	return m, err
}

func runMap(m mapa, startPos, endPos string) (numSteps int) {
	var (
		idx        int    = 0
		length     int    = len(m.instructionsLoop)
		currentPos string = startPos
	)
	for {
		lr := m.positions[currentPos]
		if m.instructionsLoop[idx] == 'L' {
			currentPos = lr[0]
		} else {
			currentPos = lr[1]
		}
		numSteps++
		idx++
		if idx == length {
			idx = 0
		}
		if currentPos == endPos {
			return numSteps
		}
	}
}

func getSteps(s []string) (numSteps int, err error) {
	var m mapa
	if m, err = parseInput(s); err != nil {
		return numSteps, err
	}
	numSteps = runMap(m, "AAA", "ZZZ")
	return numSteps, err
}

func runMap2(m mapa, startPos, endPos string) (numSteps int) {
	var (
		length     int = len(m.instructionsLoop)
		currentPos []string
		freqSteps  []int
		wg         sync.WaitGroup
	)
	// Get initial positions
	for k, _ := range m.positions {
		if string(k[len(k)-1]) == startPos {
			currentPos = append(currentPos, k)
		}
	}
	freqSteps = make([]int, len(currentPos))
	for i := 0; i < len(currentPos); i++ {
		wg.Add(1)
		go func(firstPosition string, freqStep *int) {
			defer wg.Done()
			var (
				idx      int    = 0
				position string = firstPosition
			)
			*freqStep = 0
			for {
				lr := m.positions[position]
				if m.instructionsLoop[idx] == 'L' {
					position = lr[0]
				} else {
					position = lr[1]
				}
				*freqStep++
				if string(position[len(position)-1]) == endPos {
					break
				}
				idx++
				if idx == length {
					idx = 0
				}
			}
		}(currentPos[i], &freqSteps[i])
	}
	wg.Wait()
	// Less common multiple
	numSteps = util.LCM(freqSteps[0], freqSteps[1:])
	return numSteps
}

func getSteps2(s []string) (numSteps int, err error) {
	var m mapa
	if m, err = parseInput(s); err != nil {
		return numSteps, err
	}
	numSteps = runMap2(m, "A", "Z")
	return numSteps, err
}

func main() {
	// abs, _ := filepath.Abs("../input.txt")
	// output, _ := file.ReadInput(abs)
	lines := util.GetFile("../../test.txt")
	//fmt.Println(getSteps(lines))
	fmt.Println(getSteps2(lines)) // 15995167053923
}
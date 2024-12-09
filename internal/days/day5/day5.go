package day5

import (
	"slices"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	rules, lines, linesRev := parse(input)

	print("Rules: %+v\n", rules)
	print("Lines: %+v\n", lines)
	print("LinesRev: %+v\n", linesRev)

	var acc int

	// each line
	for i, lineRev := range linesRev {
		// check if line follows rules
		if checkLine(lineRev, rules) {
			acc += lines[i][len(lines[i])/2]
		}
	}

	return strconv.Itoa(acc)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	rules, lines, linesRev := parse(input)

	print("Rules: %+v\n", rules)
	print("Lines: %+v\n", lines)

	var acc int

	// each line
	for i, lineRev := range linesRev {
		// check if line follows rules
		if checkLine(lineRev, rules) {
			continue
		}

		line := lines[i]

		// construct line to follow rule
		for {
			for i := 1; i < len(line); i++ {
				num := line[i]
				// check for rules for current num
				otherNums, ok := rules[num]
				if !ok {
					continue
				}
				// find one of numbers that should be after current
				for _, otherNum := range otherNums {
					idx := slices.Index(line[:i], otherNum)
					if idx == -1 {
						continue
					}
					// put current before found number
					move(line, i, idx)
					// move to next number in line
					break
				}
			}

			if checkLine(revLine(line), rules) {
				break
			}
		}

		acc += line[len(line)/2]
	}

	return strconv.Itoa(acc)
}

func checkLine(line map[int]int, rules map[int][]int) bool {
	for num, idx := range line {
		// check if there is rule for num
		otherNums, ok := rules[num]
		if !ok {
			continue
		}

		for _, otherNum := range otherNums {
			// check if there is otherNum in line
			otherIdx, ok := line[otherNum]
			if !ok {
				continue
			}

			if idx > otherIdx {
				return false
			}
		}
	}

	return true
}

func parse(input []string) (rules map[int][]int, lines [][]int, linesRev []map[int]int) {
	input = utils.FilterEmptyLines(input)
	dataLine := slices.IndexFunc(input, func(s string) bool {
		return s[2] == ','
	})
	rulesSlice := input[:dataLine]
	linesSlice := input[dataLine:]

	rules = make(map[int][]int, len(rulesSlice))
	for i := range rulesSlice {
		rules[parseInt(rulesSlice[i][:2])] = append(rules[parseInt(rulesSlice[i][:2])], parseInt(rulesSlice[i][3:]))
	}

	lines = make([][]int, 0, len(linesSlice))
	for i := range linesSlice {
		line := make([]int, len(linesSlice[i])/3+1)
		for j := 0; j < len(linesSlice[i])/3+1; j++ {
			num := parseInt(linesSlice[i][j*3 : j*3+2])
			line[j] = num
		}
		linesRev = append(linesRev, revLine(line))
		lines = append(lines, line)
	}

	return
}

func revLine(line []int) map[int]int {
	lineRev := make(map[int]int, len(line))
	for i, x := range line {
		lineRev[x] = i
	}
	return lineRev
}

func parseInt(str string) int {
	return int(str[0]-'0')*10 + int(str[1]-'0')
}

func move(ar []int, from, to int) []int {
	acc := ar[from]
	for i := from; i > to; i-- {
		ar[i] = ar[i-1]
	}
	ar[to] = acc
	return ar
}

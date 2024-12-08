package day2

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2024/pkg/models"
	utils2 "github.com/wmuga/aoc2024/pkg/utils"
)

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug bool, canErr bool) string {
	print := utils.DebugPrint(debug)
	data := parse(input)
	print("Parsed %d lines\n", len(data))

	acc := 0

	for _, line := range data {
		if checkLine(models.PrintFunc(print), line, 0, -1, canErr) {
			acc++
		}
	}

	return strconv.Itoa(acc)
}

func checkLine(print models.PrintFunc, line []int, start, skip int, canErr bool) bool {
	start = max(0, start)

	print("Line: %v. skip: %d\n", line, skip)

	for i := start; i < len(line)-2; i++ {
		if i == skip {
			continue
		}

		j := i + 1
		if j == skip {
			j++
		}

		if j >= len(line) {
			break
		}

		k := j + 1
		if k == skip {
			k++
		}

		if k >= len(line) {
			break
		}

		diff1 := line[i] - line[j]
		diff2 := line[j] - line[k]

		// localmin or localmax
		if (diff1 > 0 && diff2 < 0) || (diff1 < 0 && diff2 > 0) {
			print("Wrong diff: Line: %v. diff1: %d; diff2: %d; i: %d; j: %d; k: %d", line, diff1, diff2, i, j, k)
			// Cant skip
			if !canErr {
				return false
			}
			// try to skip one of 3
			return checkLine(print, line, i-2, i, false) || checkLine(print, line, i-2, j, false) || checkLine(print, line, i-2, k, false)
		}

		diff1 = utils.Abs(diff1)
		diff2 = utils.Abs(diff2)

		if 1 > diff1 || diff1 > 3 || 1 > diff2 || diff2 > 3 {
			print("Big diff: Line: %v. diff1: %d; diff2: %d; i: %d; j: %d; k: %d", line, diff1, diff2, i, j, k)
			// Cant skip
			if !canErr {
				return false
			}
			// try to skip one of 3
			return checkLine(print, line, i-2, i, false) || checkLine(print, line, i-2, j, false) || checkLine(print, line, i-2, k, false)
		}
	}

	return true
}

func parse(input []string) [][]int {
	input = utils.FilterEmptyLines(input)
	data := make([][]int, len(input))
	for i := range input {
		ints, err := utils2.ParseIntsLine(input[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data[i] = ints
	}
	return data
}

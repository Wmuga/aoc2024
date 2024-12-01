package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	intsL, intsR := parse(input)
	print("Parsed: %d lines\n", len(intsL))

	sort.SliceStable(intsL, func(i, j int) bool {
		return intsL[i] < intsL[j]
	})

	sort.SliceStable(intsR, func(i, j int) bool {
		return intsR[i] < intsR[j]
	})

	var acc int

	for i := range len(intsL) {
		acc += utils.Abs(intsL[i] - intsR[i])
	}

	return strconv.Itoa(acc)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	intsL, intsR := parse(input)
	print("Parsed: %d lines\n", len(intsL))

	counts := map[int]int{}
	for _, num := range intsR {
		counts[num]++
	}

	var acc int
	for _, num := range intsL {
		acc += num * counts[num]
	}

	return strconv.Itoa(acc)
}

func parse(input []string) ([]int, []int) {
	input = utils.FilterEmptyLines(input)
	resL := make([]int, 0, len(input))
	resR := make([]int, 0, len(input))
	for _, line := range input {
		ints := strings.Split(line, "   ")
		if len(ints) != 2 {
			fmt.Println("Malformed input:", line)
			os.Exit(1)
		}
		resL = append(resL, utils.Must(strconv.Atoi(ints[0])))
		resR = append(resR, utils.Must(strconv.Atoi(ints[1])))
	}
	return resL, resR
}

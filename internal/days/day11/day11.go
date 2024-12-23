package day11

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
	utils2 "github.com/wmuga/aoc2024/pkg/utils"
)

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, 25)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, 75)
}

func solve(input []string, debug bool, stepCount int) string {
	print := utils.DebugPrint(debug)

	data := parse(input)
	print("Parsed input. Stones: %v\n", data)

	for i := 0; i < stepCount; i++ {
		data2 := map[uint64]uint64{}

		for num, count := range data {
			// 0 replaced by 1
			if num == 0 {
				data2[1] += count
				continue
			}

			// even number of digits. split into two stones
			numStr := strconv.FormatUint(num, 10)
			if len(numStr)%2 == 0 {
				num1, err := strconv.ParseUint(numStr[:len(numStr)/2], 10, 64)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				num2, err := strconv.ParseUint(numStr[len(numStr)/2:], 10, 64)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				data2[num1] += count
				data2[num2] += count
				continue
			}

			// multiply by 2024
			data2[num*2024] += count
		}

		data = data2

		if debug {
			print("Step %d; Count:%d\n", i+1, sumStones(data))
		}
	}

	return strconv.FormatUint(sumStones(data), 10)
}

func sumStones(m map[uint64]uint64) uint64 {
	var acc uint64
	for _, v := range m {
		acc += v
	}
	return acc
}

func parse(input []string) map[uint64]uint64 {
	input = utils.FilterEmptyLines(input)
	stones, err := utils2.GetInts(input[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	grouped := map[uint64]uint64{}
	for _, num := range stones {
		grouped[uint64(num)]++
	}

	return grouped
}

package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

var (
	reMulOp = regexp.MustCompile(`(mul\([+-]?\d+,[+-]?\d+\))|(do\(\))|(don't\(\))`)
)

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, true)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	_ = print
	return solve(input, debug, false)
}

func solve(input []string, debug bool, skipDo bool) string {
	print := utils.DebugPrint(debug)

	allStrs := reMulOp.FindAllString(strings.Join(input, "\n"), -1)
	do := true

	print("Found %d valid ops\n", len(allStrs))

	var acc int64
	for _, op := range allStrs {
		switch op {
		case "don't()":
			if !skipDo {
				do = false
			}
		case "do()":
			if !skipDo {
				do = true
			}
			continue
		default:
			if do {
				acc += mul(op)
			}
		}
	}

	return strconv.FormatInt(acc, 10)
}

func mul(mulstr string) int64 {
	// mul(X,Y)
	idxComma := strings.Index(mulstr, ",")
	num1Str := mulstr[4:idxComma]
	num2Str := mulstr[idxComma+1 : len(mulstr)-1]

	num1 := utils.Must(strconv.ParseInt(num1Str, 10, 64))
	num2 := utils.Must(strconv.ParseInt(num2Str, 10, 64))

	return num1 * num2
}

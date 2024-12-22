package day7

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
	utils2 "github.com/wmuga/aoc2024/pkg/utils"
)

type Day struct{}

type op int

const (
	opAdd op = iota
	opMul
	opConcat
)

type equasion struct {
	res  int64
	nums []int64
	ops  []op
}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug, day2 bool) string {
	print := utils.DebugPrint(debug)
	eqs := parse(input)
	print("Equasion count: %d\n", len(eqs))

	var acc int64
	for _, eq := range eqs {
		if canSolveReq(eq.nums[0], 0, eq, day2) {
			acc += eq.res
		}
	}

	return strconv.FormatInt(acc, 10)
}

func canSolveReq(acc int64, idx int, e *equasion, day2 bool) bool {
	if idx >= len(e.ops) {
		return acc == e.res
	}

	r := canSolveReq(acc+e.nums[idx+1], idx+1, e, day2)
	if r {
		return true
	}
	e.ops[idx] = opMul
	r = canSolveReq(acc*e.nums[idx+1], idx+1, e, day2)
	if r || !day2 {
		return r
	}

	e.ops[idx] = opConcat
	return canSolveReq(concat(acc, e.nums[idx+1]), idx+1, e, day2)
}

func concat(l, r int64) int64 {
	concatStr := strconv.FormatInt(l, 10) + strconv.FormatInt(r, 10)
	concatRes, _ := strconv.ParseInt(concatStr, 10, 64)
	return concatRes
}

func parse(input []string) []*equasion {
	res := make([]*equasion, 0, len(input))
	for _, line := range input {
		if line == "" {
			continue
		}

		nums, err := utils2.GetInts(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if len(nums) < 2 {
			fmt.Fprintln(os.Stderr, "Not enough ints", nums)
			os.Exit(1)
		}

		res = append(res, &equasion{
			res:  nums[0],
			nums: nums[1:],
			ops:  make([]op, len(nums)-2),
		})
	}
	return res
}

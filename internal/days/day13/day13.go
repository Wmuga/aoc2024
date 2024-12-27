package day13

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"

	utils2 "github.com/wmuga/aoc2024/pkg/utils"
)

type Day struct{}

const (
	maxPressCount = 100

	aTokens = 3
	bTokens = 1
)

type point struct {
	X int64
	Y int64
}

type game struct {
	btnA  point
	btnB  point
	prize point
}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug, part2 bool) string {
	print := utils.DebugPrint(debug)
	games := parse(input)

	print("Parsed input. Games count: %d\n", len(games))

	var acc int64
	for i := range games {
		if part2 {
			games[i].prize.X += 10000000000000
			games[i].prize.Y += 10000000000000
		}

		tokens := calcTokens(games[i], part2)
		if tokens == 0 {
			print("Game %d is unplayable", i+1)
			continue
		}
		acc += tokens
	}

	return strconv.FormatInt(acc, 10)
}

func calcTokens(g game, uncapped bool) int64 {
	if !uncapped && ((g.btnA.X+g.btnB.X)*maxPressCount < g.prize.X ||
		(g.btnA.Y+g.btnB.Y)*maxPressCount < g.prize.Y) {
		return 0
	}

	b := (g.btnA.X*g.prize.Y - g.prize.X*g.btnA.Y) / (g.btnA.X*g.btnB.Y - g.btnB.X*g.btnA.Y)

	if b < 0 || (!uncapped && b > maxPressCount) {
		return 0
	}

	a := (g.prize.X - g.btnB.X*b) / g.btnA.X
	if a < 0 || (!uncapped && a > maxPressCount) {
		return 0
	}

	if g.btnA.X*a+g.btnB.X*b != g.prize.X ||
		g.btnA.Y*a+g.btnB.Y*b != g.prize.Y {
		return 0
	}

	return a*aTokens + b*bTokens
}

func parse(input []string) []game {
	input = utils.FilterEmptyLines(input)

	if len(input)%3 != 0 {
		fmt.Fprintln(os.Stderr, "wrong lines count", len(input))
		os.Exit(1)
	}

	games := make([]game, len(input)/3)
	for i := 0; i < len(games); i++ {
		offset := i * 3
		xA, yA := mustTwoInts(input[offset+0])
		xB, yB := mustTwoInts(input[offset+1])
		xP, yP := mustTwoInts(input[offset+2])
		games[i] = game{
			btnA:  point{X: xA, Y: yA},
			btnB:  point{X: xB, Y: yB},
			prize: point{X: xP, Y: yP},
		}
	}

	return games
}

func mustTwoInts(line string) (int64, int64) {
	int1, idx, err := utils2.NextInt(line, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	int2, _, err := utils2.NextInt(line, idx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return int1, int2
}

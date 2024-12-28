package day14

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2024/pkg/models"
	utils2 "github.com/wmuga/aoc2024/pkg/utils"
)

type point = models.Point2D

type robot struct {
	pos point
	vel point
}

func (r robot) posAt(step int, field point) point {
	return point{
		X: (r.pos.X + r.vel.X*step) % field.X,
		Y: (r.pos.Y + r.vel.Y*step) % field.Y,
	}
}

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	fieldSize := point{X: 101, Y: 103}

	if len(input) < 20 {
		fieldSize = point{X: 11, Y: 7}
	}

	middle := point{X: fieldSize.X / 2, Y: fieldSize.Y / 2}

	robots := parse(input, fieldSize)
	print("Parsed %d robots\n", len(robots))

	qs := [4]int{}

	for i := range robots {
		pos := robots[i].posAt(100, fieldSize)

		if pos.X < middle.X && pos.Y < middle.Y {
			qs[0]++
			continue
		}
		if pos.X > middle.X && pos.Y < middle.Y {
			qs[1]++
			continue
		}
		if pos.X < middle.X && pos.Y > middle.Y {
			qs[2]++
			continue
		}
		if pos.X > middle.X && pos.Y > middle.Y {
			qs[3]++
			continue
		}
	}

	return strconv.FormatInt(int64(qs[0]*qs[1]*qs[2]*qs[3]), 10)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	fieldSize := point{X: 101, Y: 103}

	robots := parse(input, fieldSize)
	print("Parsed %d robots\n", len(robots))

	// assuming, when tree forms there are rows and columns with higher robot count
	var (
		maxX, maxY   int
		stepX, stepY int
	)

	// cycle repeats by modulo of ring. Checking for step with most
	for step := 1; step < 103; step++ {
		byX := [101]int{}
		byY := [103]int{}
		for i := range robots {
			pos := robots[i].posAt(step, fieldSize)
			byX[pos.X]++
			byY[pos.Y]++
		}

		curMaxX := slices.Max(byX[:])
		curMaxY := slices.Max(byY[:])

		if curMaxX > maxX {
			maxX = curMaxX
			stepX = step
		}

		if curMaxY > maxY {
			maxY = curMaxY
			stepY = step
		}
	}

	res := utils2.CRT(utils2.CRTArg{Rem: stepX, Mod: fieldSize.X}, utils2.CRTArg{Rem: stepY, Mod: fieldSize.Y})

	return strconv.FormatInt(int64(res), 10)
}

func parse(input []string, fieldSize point) []robot {
	input = utils.FilterEmptyLines(input)

	robots := make([]robot, len(input))
	for i := range input {
		x, idx, err := utils2.NextInt(input[i], 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		y, idx, err := utils2.NextInt(input[i], idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		vx, idx, err := utils2.NextInt(input[i], idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		vy, _, err := utils2.NextInt(input[i], idx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		robots[i] = robot{
			pos: point{X: int(x), Y: int(y)},
			// make additive inverse from negative numbers
			vel: inverse(point{X: int(vx), Y: int(vy)}, fieldSize),
		}
	}

	return robots
}

func inverse(p point, ring point) point {
	if p.X < 0 {
		p.X = (p.X + (-p.X/ring.X+1)*ring.X) % ring.X
	}
	if p.Y < 0 {
		p.Y = (p.Y + (-p.Y/ring.Y+1)*ring.Y) % ring.Y
	}
	return p
}

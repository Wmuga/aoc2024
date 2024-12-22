package day8

import (
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"

	"github.com/wmuga/aoc2024/pkg/models"
	"github.com/wmuga/aoc2024/pkg/set"
)

type point = models.Point2D

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug, day2 bool) string {
	print := utils.DebugPrint(debug)

	antennas, size := parse(input)

	print("Parsed input. Distinct antennas freq: %d. Field size: %s\n", len(antennas), size)

	antiNodes := set.New[point]()

	for _, ants := range antennas {
		for i := 0; i < len(ants)-1; i++ {
			for j := i + 1; j < len(ants); j++ {
				// get antennas
				a1 := ants[i]
				a2 := ants[j]

				// get vector
				vec := a2.Sub(a1)

				// get position of antinodes
				pos1 := a1.Sub(vec)
				pos2 := a2.Add(vec)

				// add antinodes if inbound
				if pos1.InBoundOf(size) {
					antiNodes.Upsert(pos1)
				}

				if pos2.InBoundOf(size) {
					antiNodes.Upsert(pos2)
				}

				if !day2 {
					continue
				}
				// add antennas as nodes
				antiNodes.Upsert(a1)
				antiNodes.Upsert(a2)

				// while inbound. add more
				for pos1.InBoundOf(size) {
					antiNodes.Upsert(pos1)
					pos1 = pos1.Sub(vec)
				}

				for pos2.InBoundOf(size) {
					antiNodes.Upsert(pos2)
					pos2 = pos2.Add(vec)
				}
			}
		}
	}

	return strconv.Itoa(antiNodes.Len())
}

func parse(input []string) (antennas map[rune][]point, size point) {
	input = utils.FilterEmptyLines(input)

	size = point{X: len(input[0]), Y: len(input)}
	antennas = make(map[rune][]point)

	for y, line := range input {
		for x, c := range line {
			if c == '.' {
				continue
			}

			antennas[c] = append(antennas[c], point{X: x, Y: y})
		}
	}

	return
}

package day10

import (
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"

	"github.com/wmuga/aoc2024/pkg/models"
	"github.com/wmuga/aoc2024/pkg/set"
)

type point = models.Point2D

type field struct {
	zeros   []point
	heights [][]int
}

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug, part2 bool) string {
	print := utils.DebugPrint(debug)

	f := parse(input)

	print("Parsed field: %d:%d; Trails: %d", len(f.heights[0]), len(f.heights), len(f.zeros))

	// caching for point to heights
	cache := map[point][]point{}
	var acc int
	for _, start := range f.zeros {
		pts := calcTrailRec(start, start, f.heights, cache)
		var val int
		// part 2 - distinct paths
		if part2 {
			val = len(pts)
		} else {
			// part 1 - distinct heights
			s := set.New[point]()
			for _, pt := range pts {
				s.Upsert(pt)
			}
			val = s.Len()
		}

		print("Trail %s; val: %d", start, val)
		acc += val
	}

	return strconv.Itoa(acc)
}

func calcTrailRec(start, cur point, heights [][]int, cache map[point][]point) []point {
	if val, ok := cache[cur]; ok {
		return val
	}

	if heights[cur.Y][cur.X] == 9 {
		return []point{cur}
	}

	var acc []point
	// scan nearby heights
	if cur.X > 0 && heights[cur.Y][cur.X-1]-heights[cur.Y][cur.X] == 1 {
		pts := calcTrailRec(start, point{X: cur.X - 1, Y: cur.Y}, heights, cache)
		acc = pAppend(acc, pts)
	}
	if cur.X < len(heights[cur.Y])-1 && heights[cur.Y][cur.X+1]-heights[cur.Y][cur.X] == 1 {
		pts := calcTrailRec(start, point{X: cur.X + 1, Y: cur.Y}, heights, cache)
		acc = pAppend(acc, pts)
	}
	if cur.Y > 0 && heights[cur.Y-1][cur.X]-heights[cur.Y][cur.X] == 1 {
		pts := calcTrailRec(start, point{X: cur.X, Y: cur.Y - 1}, heights, cache)
		acc = pAppend(acc, pts)
	}
	if cur.Y < len(heights)-1 && heights[cur.Y+1][cur.X]-heights[cur.Y][cur.X] == 1 {
		pts := calcTrailRec(start, point{X: cur.X, Y: cur.Y + 1}, heights, cache)
		acc = pAppend(acc, pts)
	}

	cache[cur] = acc
	return acc
}

func pAppend(ar1, ar2 []point) []point {
	if len(ar1) == 0 {
		return ar2
	}

	if len(ar2) == 0 {
		return ar1
	}

	arRes := make([]point, len(ar1)+len(ar2))
	copy(arRes, ar1)
	copy(arRes[len(ar1):], ar2)
	return arRes
}

func parse(input []string) field {
	input = utils.FilterEmptyLines(input)

	heights := make([][]int, len(input))
	zeros := make([]point, 0)

	for y, line := range input {
		heights[y] = make([]int, len(line))
		for x, c := range line {
			h := int(c - '0')
			heights[y][x] = h
			if h == 0 {
				zeros = append(zeros, point{X: x, Y: y})
			}
		}
	}

	return field{
		heights: heights,
		zeros:   zeros,
	}
}

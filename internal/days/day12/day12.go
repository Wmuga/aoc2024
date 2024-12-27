package day12

import (
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"

	"github.com/wmuga/aoc2024/pkg/models"
	"github.com/wmuga/aoc2024/pkg/set"
)

type point = models.Point2D

type position struct {
	p   point
	dir int
}

type fieldPart struct {
	letter    rune
	cells     *set.Set[point]
	perimeter int
}

type dirSlice []point

// up, right, down, left
var dirs = dirSlice{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

func (d dirSlice) AddNext(p point, i int) point {
	return p.Add(d[(i+1)%len(d)])
}

func (d dirSlice) AddPrev(p point, i int) point {
	return p.Add(d[(i-1+len(d))%len(d)])
}

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug bool, part2 bool) string {
	print := utils.DebugPrint(debug)

	field := parse(input)
	print("Parsed field %d:%d\n", len(field[0]), len(field))

	var acc int64

	visitedAll := set.New[point]()
	for y, line := range field {
		for x, letter := range line {
			p := point{X: x, Y: y}
			// check if already visited
			if visitedAll.Contains(p) {
				continue
			}

			f := fieldPart{
				letter: letter,
				cells:  set.New[point](),
			}
			parseFieldRec(p, field, &f, visitedAll)

			if !part2 {
				acc += int64(f.cells.Len()) * int64(f.perimeter)
				print("Field: %c; Start: %s; Perimeter: %d; Area: %d\n", f.letter, p, f.perimeter, f.cells.Len())
				continue
			}

			sidesCount := calcSides(f)
			acc += int64(f.cells.Len()) * int64(sidesCount)
			print("Field: %c; Start: %s; Sides: %d; Area: %d\n", f.letter, p, sidesCount, f.cells.Len())
		}
	}

	return strconv.FormatInt(acc, 10)
}

func calcSides(p fieldPart) int {
	sides := 0

	visited := set.New[point]()
outer:
	for cur := range p.cells.Iterator() {
		if visited.Contains(cur) {
			continue
		}

		visited.Upsert(cur)

		for i := 0; i < len(dirs); i++ {
			if !p.cells.Contains(cur.Add(dirs[i])) {
				pos := position{cur, (i + 1) % len(dirs)}
				sides += calcSide(pos, p, visited)
				continue outer
			}
		}
	}

	return sides
}

func calcSide(start position, p fieldPart, visited *set.Set[point]) int {
	count := 0
	cur := start

	// circle around edges until
	for i := 0; i == 0 || start != cur; i++ {
		visited.Upsert(cur.p)

		// is there cell on top of current
		if p.cells.Contains(dirs.AddPrev(cur.p, cur.dir)) {
			count++
			cur.p = dirs.AddPrev(cur.p, cur.dir)
			cur.dir = (cur.dir - 1 + len(dirs)) % len(dirs)
			continue
		}

		// is there cell in front of current
		if p.cells.Contains(cur.p.Add(dirs[cur.dir])) {
			cur.p = cur.p.Add(dirs[cur.dir])
			continue
		}

		// turn around same point
		count++
		cur.dir = (cur.dir + 1) % len(dirs)
	}

	return count
}

func parseFieldRec(cur point, field [][]rune, p *fieldPart, global *set.Set[point]) {
	// Check for out of bound
	if cur.X < 0 || cur.Y < 0 || cur.Y >= len(field) || cur.X >= len(field[cur.Y]) {
		p.perimeter++
		return
	}

	// check for same letter
	if p.letter != field[cur.Y][cur.X] {
		p.perimeter++
		return
	}

	// check if already visited
	if global.Contains(cur) {
		return
	}

	// set point as visited
	global.Upsert(cur)
	p.cells.Upsert(cur)

	// check neighbours
	parseFieldRec(cur.Add(dirs[0]), field, p, global)
	parseFieldRec(cur.Add(dirs[1]), field, p, global)
	parseFieldRec(cur.Add(dirs[2]), field, p, global)
	parseFieldRec(cur.Add(dirs[3]), field, p, global)
}

func parse(input []string) [][]rune {
	input = utils.FilterEmptyLines(input)

	res := make([][]rune, len(input))
	for i := range input {
		res[i] = []rune(input[i])
	}

	return res
}

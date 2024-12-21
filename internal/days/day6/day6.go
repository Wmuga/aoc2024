package day6

import (
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"

	"github.com/wmuga/aoc2024/pkg/set"
)

type point struct {
	x, y int
}

func (p point) Add(o point) point {
	return point{p.x + o.x, p.y + o.y}
}

func (p point) Sub(o point) point {
	return point{p.x - o.x, p.y - o.y}
}

func (p point) String() string {
	return "{" + strconv.Itoa(p.x) + "; " + strconv.Itoa(p.y) + "}"
}

type field struct {
	startPos point
	walls    *set.Set[point]
	size     point
}

var dirs = []point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	f := parse(input)

	print("Start position: %s\nWall count: %d;\nField size: %s\n", f.startPos, f.walls.Len(), f.size)

	stepped := calcPath(debug, f)

	return strconv.Itoa(stepped.Len())
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)
	f := parse(input)

	print("Start position: %s\nWall count: %d;\nField size: %s\n", f.startPos, f.walls.Len(), f.size)

	byX, byY := groupWalls(f.walls)

	counter := 0
	path := calcPath(debug, f)
	for pos := range path.Iterator() {
		// can't place on guard
		if pos == f.startPos {
			continue
		}

		byX2 := addCloned(byX, pos.x, pos.y)
		byY2 := addCloned(byY, pos.y, pos.x)

		if checkLoop(f.startPos, byX2, byY2) {
			counter++
		}
	}

	return strconv.Itoa(counter)
}

func checkLoop(start point, byX, byY map[int][]int) bool {
	dir := 0
	type visit struct {
		d int
		p point
	}
	pos := start
	visited := set.New[visit]()
	for {
		v := visit{dir, pos}
		if _, ok := visited.Get(v); ok {
			return true
		}

		visited.Upsert(v)

		var nextCoord = -1
		switch dir {
		case 0:
			if len(byX[pos.x]) == 0 {
				return false
			}
			nextCoord = minCoord(byX[pos.x], pos.y, nextCoord, false)
			pos.y = nextCoord + 1
		case 1:
			if len(byY[pos.y]) == 0 {
				return false
			}
			nextCoord = minCoord(byY[pos.y], pos.x, nextCoord, true)
			pos.x = nextCoord - 1
		case 2:
			if len(byX[pos.x]) == 0 {
				return false
			}
			nextCoord = minCoord(byX[pos.x], pos.y, nextCoord, true)
			pos.y = nextCoord - 1
		default:
			if len(byY[pos.y]) == 0 {
				return false
			}
			nextCoord = minCoord(byY[pos.y], pos.x, nextCoord, false)
			pos.x = nextCoord + 1
		}

		if nextCoord == -1 {
			return false
		}

		dir = (dir + 1) % len(dirs)
	}
}

func minCoord(m []int, start, maxVal int, asc bool) int {
	minDiff := 0
	res := maxVal

	for _, v := range m {
		diff := v - start
		if asc && diff > 0 && (minDiff == 0 || diff < minDiff) {
			res = v
			minDiff = diff
			continue
		}

		if !asc && diff < 0 && (minDiff == 0 || -diff < minDiff) {
			res = v
			minDiff = -diff
		}
	}

	return res
}

func calcPath(debug bool, f field) *set.Set[point] {
	print := utils.DebugPrint(debug)
	stepped := set.New[point]()
	pos := f.startPos
	dir := 0

	for {
		// step in direction
		stepped.Upsert(pos)
		newPos := pos.Add(dirs[dir])
		// check for out of bounds
		if newPos.x < 0 || newPos.y < 0 || newPos.x >= f.size.x || newPos.y >= f.size.y {
			break
		}
		// check for wall
		if _, ok := f.walls.Get(newPos); ok {
			print("Wall at %s\n", newPos)
			dir = (dir + 1) % len(dirs)
			continue
		}
		// apply step
		pos = newPos
	}
	return stepped
}

func parse(input []string) (f field) {
	f.size = point{x: len(input[0]), y: len(input)}
	f.walls = set.New[point]()
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			switch input[y][x] {
			case '.':
				continue
			case '^':
				f.startPos = point{x, y}
			case '#':
				f.walls.Upsert(point{x, y})
			}
		}
	}
	return f
}

func groupWalls(walls *set.Set[point]) (byX map[int][]int, byY map[int][]int) {
	byX = make(map[int][]int)
	byY = make(map[int][]int)
	for wall := range walls.Iterator() {
		byX[wall.x] = append(byX[wall.x], wall.y)
		byY[wall.y] = append(byY[wall.y], wall.x)
	}

	return
}

func addCloned(to map[int][]int, key, value int) map[int][]int {
	res := make(map[int][]int, len(to))
	for k, v := range to {
		vCloned := make([]int, len(v))
		copy(vCloned, v)
		res[k] = v
	}
	res[key] = append(res[key], value)
	return res
}

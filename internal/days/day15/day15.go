package day15

import (
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
	"github.com/wmuga/aoc2024/pkg/models"
)

type point = models.Point2D

var dirs = []point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

func moveToVec(move rune) point {
	switch move {
	case 'v':
		return dirs[2]
	case '^':
		return dirs[0]
	case '<':
		return dirs[3]
	case '>':
		return dirs[1]
	}
	return point{}
}

type Day struct{}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	print := utils.DebugPrint(debug)

	parse(input)

	_ = print
	return ""
}

func solve(input []string, debug, part2 bool) string {
	print := utils.DebugPrint(debug)

	field, pos, moves := parse(input)
	var plusHalf bool
	_ = plusHalf
	print("Parsed field %d;%d. Start at %s. Moves lines: %d", len(field[0]), len(field), pos, len(moves))

	for _, move := range moves {
		vec := moveToVec(move)
		newPos := pos.Add(vec)
		switch field[newPos.Y][newPos.X] {
		// can't move walls
		case '#':
			continue
		// try to move barrel
		case 'O', '0':
		searchloop:
			for barrelPos := newPos.Add(vec); ; barrelPos = barrelPos.Add(vec) {
				switch field[barrelPos.Y][barrelPos.X] {
				// place to put a barrel
				case '.':
					// part 1 or vertically barrel moves with full coordinates
					if !part2 || vec == dirs[0] || vec == dirs[2] {
						field[barrelPos.Y][barrelPos.X] = field[newPos.Y][newPos.X]
						field[newPos.Y][newPos.X] = '.'
						pos = newPos
					}
					break searchloop
				// can't place on wall
				case '#':
					break searchloop
				// can move multiple barrels
				default:
					continue
				}
			}
		default:
			pos = newPos
		}
	}

	var acc int64
	for y, line := range field {
		for x, c := range line {
			if c == 'O' {
				if !part2 {
					acc += int64(y*100 + x)
				} else {
					acc += int64(y*100 + x*2)
				}
				continue
			}
			if c == '0' {
				acc += int64(y*100 + x*2 + 1)
			}
		}
	}

	if debug {
		print("Field:\n%s\n", strings.Join(fieldToString(field, pos), "\n"))
	}

	return strconv.FormatInt(acc, 10)
}

func parse(input []string) (field [][]rune, start point, moves []rune) {
	field = make([][]rune, 0, len(input))
	var y int

	for y = range input {
		if input[y] == "" || input[y][0] != '#' {
			break
		}
		field = append(field, []rune(input[y]))
		x := strings.IndexByte(input[y], '@')
		if x == -1 {
			continue
		}
		start = point{X: x, Y: y}
		field[y][x] = '.'
	}

	moves = make([]rune, 0, len(input[y]))
	for ; y < len(input); y++ {
		moves = append(moves, []rune(input[y])...)
	}

	return
}

func fieldToString(s [][]rune, pos point) []string {
	res := make([]string, len(s))
	for i := range s {
		if i == pos.Y {
			res[i] = string(s[i][:pos.X]) + "@" + string(s[i][pos.X+1:])
			continue
		}
		res[i] = string(s[i])
	}
	return res
}

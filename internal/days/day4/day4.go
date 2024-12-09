package day4

import (
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

type vec2 struct {
	x int
	y int
}

const (
	wordXMAS = "XMAS"
	wordMAS  = "MAS"
)

var steps = []vec2{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug, part2 bool) string {
	print := utils.DebugPrint(debug)
	data := parse(input)
	print("Input: %d:%d chars\n", len(data[0]), len(data))

	var acc int

	calcAcc := accPart1
	if part2 {
		calcAcc = accPart2
	}

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			acc += calcAcc(data, y, x, print)
			continue
		}
	}

	return strconv.Itoa(acc)
}

// accPart1 calculate accumulator with part 1 logic
// data - input characters
// y, x - current point
// print - debug print function
func accPart1(data [][]rune, y, x int, print models.PrintFunc) int {
	var acc int

	for i := 0; i < len(steps); i++ {
		if chechLine(data, y, x, steps[i].y, steps[i].x, wordXMAS) {
			print("Found at: %d:%d with vec {%d, %d}\n", x, y, steps[i].x, steps[i].y)
			acc++
		}
	}

	return acc
}

// accPart2 calculate accumulator with part 2 logic
// data - input characters
// y, x - current point
// print - debug print function
func accPart2(data [][]rune, y, x int, print models.PrintFunc) int {
	// Can't check if not enough data for cross from top left corner
	if y+2 >= len(data) || x+2 >= len(data[y]) {
		return 0
	}

	// TL - DR
	if chechLine(data, y, x, 1, 1, wordMAS) {
		// DL - TR or TR - DL
		if chechLine(data, y+2, x, -1, 1, wordMAS) || chechLine(data, y, x+2, 1, -1, wordMAS) {
			print("Found at: %d:%d with vec {-1,  1}\n", x, y)
			return 1
		}
		return 0
	}

	// DL - TR and DR - TL
	if chechLine(data, y+2, x, -1, 1, wordMAS) && chechLine(data, y+2, x+2, -1, -1, wordMAS) {
		print("Found at: %d:%d with vec {-1, 1} and {-1, -1}\n", x, y)
		return 1
	}

	// DR - TL and TR - DL
	if chechLine(data, y+2, x+2, -1, -1, wordMAS) && chechLine(data, y, x+2, 1, -1, wordMAS) {
		print("Found at: %d:%d with vec {-1, -1} and { 1, -1}\n", x, y)
		return 1
	}

	return 0
}

// chechLine - check for [word] in [input] on line
// input - input chars
// y,x - start point
// dY, dX - steps on each coordinates
// word - word to compare to
func chechLine(input [][]rune, y, x, dY, dX int, word string) bool {
	stepsCount := len(word)

	lastX := x + dX*(stepsCount-1)
	lastY := y + dY*(stepsCount-1)

	if lastX < 0 || lastX >= len(input[y]) {
		return false
	}

	if lastY < 0 || lastY >= len(input) {
		return false
	}

	builder := strings.Builder{}
	builder.Grow(stepsCount)

	i := y
	j := x
	for step := 0; step < stepsCount; step++ {
		builder.WriteRune(input[i][j])
		i += dY
		j += dX
	}

	return builder.String() == word
}

func parse(input []string) [][]rune {
	res := make([][]rune, len(input))
	for i := range input {
		res[i] = []rune(input[i])
	}
	return res
}

package day9

import (
	"strconv"

	"github.com/wmuga/aoc2019/pkg/utils"
)

type Day struct{}

type block struct {
	id    int
	count int
}

func (Day) Solve1(input []string, debug bool) string {
	return solve(input, debug, false)
}

func (Day) Solve2(input []string, debug bool) string {
	return solve(input, debug, true)
}

func solve(input []string, debug, part2 bool) string {
	print := utils.DebugPrint(debug)

	blocks := parse(input)
	print("Parsed input. Count blocks: %d\n", len(blocks))

	leftPtr := 0
	rightPtr := len(blocks) - 1

	for leftPtr <= rightPtr {
		if blocks[leftPtr].count == 0 {
			leftPtr++
			continue
		}

		// if encountered non empty block from left skip
		if blocks[leftPtr].id != -1 {
			leftPtr++
			continue
		}

		// empty space in left. check if right pointer has values
		if blocks[rightPtr].id == -1 || blocks[rightPtr].count == 0 {
			rightPtr--
			continue
		}

		// 3 variants
		// easiest - l.count == r.count
		// move right block. move pointers closer
		if blocks[leftPtr].count == blocks[rightPtr].count {
			blocks[leftPtr].id = blocks[rightPtr].id
			blocks[rightPtr].id = -1
			leftPtr++
			rightPtr--
			continue
		}
		// l.count > r.count
		// add right block, move right pointer, substruct count from left
		if blocks[leftPtr].count > blocks[rightPtr].count {
			count := blocks[rightPtr].count
			blocks = append(blocks[:leftPtr], append([]block{blocks[rightPtr]}, blocks[leftPtr:]...)...)
			leftPtr++
			blocks[leftPtr].count -= count
			blocks[rightPtr+1].id = -1
			continue
		}
		// l.count < r.count
		// part1 - add only part of the block
		if !part2 {
			// add l.count from right. substract at right. move left
			blocks[leftPtr].id = blocks[rightPtr].id
			blocks[rightPtr].count -= blocks[leftPtr].count
			leftPtr++
			continue
		}

		// if part2 - search next
		for lPtr := leftPtr + 1; lPtr < rightPtr; lPtr++ {
			if blocks[lPtr].id == -1 && blocks[lPtr].count == blocks[rightPtr].count {
				blocks[lPtr].id = blocks[rightPtr].id
				blocks[rightPtr].id = -1
				break
			}
			// l.count > r.count
			// add right block, move right pointer, substruct count from left
			if blocks[lPtr].id == -1 && blocks[lPtr].count > blocks[rightPtr].count {
				count := blocks[rightPtr].count
				blocks = append(blocks[:lPtr], append([]block{blocks[rightPtr]}, blocks[lPtr:]...)...)
				lPtr++
				blocks[lPtr].count -= count
				blocks[rightPtr+1].id = -1
				break
			}
		}
		rightPtr--
	}

	// calc checksum
	var acc int64
	i := 0
	for _, block := range blocks {
		if block.id != -1 {
			acc += int64(i*block.id*2+block.id*(block.count-1)) * int64(block.count) / 2
		}
		i += block.count
	}

	return strconv.FormatInt(acc, 10)
}

func parse(input []string) []block {
	input = utils.FilterEmptyLines(input)
	line := input[0]
	res := make([]block, len(line))
	for i := 0; i <= len(line)/2; i++ {
		count := int(line[i*2] - '0')

		res[i*2] = block{
			id:    i,
			count: count,
		}

		if i*2+1 == len(line) {
			break
		}

		count = int(line[i*2+1] - '0')
		res[i*2+1] = block{
			id:    -1,
			count: count,
		}
	}

	return res
}

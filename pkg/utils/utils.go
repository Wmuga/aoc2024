package utils

import (
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
)

func ParseIntsLine(line string) ([]int, error) {
	data := strings.Fields(line)
	return utils.ParseIntLines(data)
}

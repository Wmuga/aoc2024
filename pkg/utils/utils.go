package utils

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/utils"
)

var (
	reInt = regexp.MustCompile(`([+-]?\d+)`)
)

func ParseIntsLine(line string) ([]int, error) {
	data := strings.Fields(line)
	return utils.ParseIntLines(data)
}

func GetInts(line string) ([]int64, error) {
	intsStr := reInt.FindAllString(line, -1)
	ints := make([]int64, len(intsStr))
	for i := range intsStr {
		var err error
		ints[i], err = strconv.ParseInt(intsStr[i], 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

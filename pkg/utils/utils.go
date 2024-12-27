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

// NextInt gets next int form string.
// startIdx - index to start from
// i - result.
// idx - index after num.
func NextInt(line string, startIdx int) (i int64, idx int, err error) {
	// states:  0 - [+-0-9], 1 - [0-9]
	state := 0
	var (
		numStart int
		numLen   int
	)

	for idx = startIdx; idx < len(line); idx++ {
		// check for +- in front of number
		if line[idx] == '+' {
			if state == 0 {
				numStart = idx
				numLen = 1
				state = 1
				continue
			}
			break
		}

		if line[idx] == '-' {
			if state == 0 {
				numStart = idx
				numLen = 1
				state = 1
				continue
			}
			break
		}

		// check for digits
		if line[idx] >= '0' && line[idx] <= '9' {
			if state == 0 {
				numStart = idx
				state = 1
			}
			numLen++
			continue
		}

		// if nothing - check if only found '+' or '-'
		if state == 1 {
			if numLen == 1 && (line[numStart] == '+' || line[numStart] == '-') {
				state = 0
				numLen = 0
				numStart = 0
				continue
			}
			break
		}
	}

	i, err = strconv.ParseInt(line[numStart:numStart+numLen], 10, 64)
	return
}

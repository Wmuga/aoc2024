package utils

import (
	"fmt"
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

type CRTArg struct {
	Rem int
	Mod int
}

func CRT(args ...CRTArg) int {
	if len(args) == 0 {
		return 0
	}

	M := int64(args[0].Mod)
	for i := 1; i < len(args); i++ {
		M *= int64(args[i].Mod)
	}

	var prime int64
	for i := range args {
		a := M / int64(args[i].Mod)
		x := egdc(a, int64(args[i].Rem), int64(args[i].Mod))
		prime += a * x
	}
	return int((prime + M) % M)
}

// a*x = b mod m
func egdc(a, b, m int64) int64 {
	var (
		r0, s0, t0 int64 = a, 1, 0
		r1, s1, t1 int64 = m, 0, 1
	)

	for r1 != 0 {
		q := r0 / r1
		r0, r1 = r1, r0-r1*q
		s0, s1 = s1, s0-s1*q
		t0, t1 = t1, t0-t1*q
	}

	return (s0*b + m) % m
}

func LazyDebugPrint(debug bool) func(format string, args ...func() []interface{}) {
	if !debug {
		return func(format string, args ...func() []interface{}) {}
	}

	return func(format string, args ...func() []interface{}) {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}

		if len(args) == 0 {
			fmt.Print(format)
			return
		}

		fmt.Printf(format, args[0]()...)
	}
}

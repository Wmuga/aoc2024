package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"

	fileparser "github.com/wmuga/aoc2019/pkg/fileParser"

	"github.com/wmuga/aoc2024/internal/days"
	"github.com/wmuga/aoc2024/pkg/picker"
)

var (
	strOk = color.GreenString("[OK]")
	strNo = color.RedString("[NO]")
)

type flags struct {
	num   int
	part2 bool
	test  bool
	debug bool
}

func parseFlags() (f flags) {
	flag.IntVar(&f.num, "n", 0, "Day number. 0 - all")
	flag.BoolVar(&f.part2, "p", true, "With part2")
	flag.BoolVar(&f.test, "t", false, "Run tests")
	flag.BoolVar(&f.debug, "d", false, "Debug actual input")
	flag.Usage = func() {
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	return f
}

func main() {
	f := parseFlags()

	dayPicker := picker.NewDayRunner()
	days.Populate(dayPicker)

	if f.num != 0 {
		day(dayPicker, f.num, f.part2, f.test, f.debug)
		return
	}

	start := time.Now()
	for i := 1; i <= dayPicker.CountDays(); i++ {
		day(dayPicker, i, f.part2, f.test, f.debug)
	}

	fmt.Println("\nSolved everything in", time.Since(start))
}

func day(dayPicker *picker.DayRunner, num int, part2 bool, test bool, debug bool) {
	// take correct day
	day, ok := dayPicker.GetDay(num)
	if !ok {
		fmt.Printf("Day %d not found\n", num)
		return
	}

	// get input and test data
	inFile, testFile := getFileNames(num)
	inData, err := fileparser.GetInput(inFile)
	if err != nil {
		fmt.Println("Error read input file:", err)
		return
	}
	testData, err := fileparser.ReadTests(testFile)
	if err != nil {
		fmt.Println("Error read test file:", err)
		return
	}

	doneTest := true
	// make tests
	if test {
		for _, test := range testData {
			var res string
			fmt.Println("Test", test.Name, "for part", test.Part)
			switch test.Part {
			case 1:
				res = day.Solve1(test.Data, true)
			case 2:
				if !part2 {
					fmt.Println("Skip part2")
					continue
				}
				res = day.Solve2(test.Data, true)
			default:
				fmt.Println("Skip unknown part", test.Part)
				continue
			}
			if test.Answer == "" {
				fmt.Println("Output:\n", res)
				continue
			}

			out := strOk
			if test.Answer != res {
				out = strNo
				doneTest = false
			}

			fmt.Printf("%s expected: %s, result: %s\n", out, test.Answer, res)
		}
	}

	if !doneTest {
		return
	}

	fmt.Println("\nDay", num, "\nAnswers:")

	start := time.Now()
	res := day.Solve1(inData, debug)
	fmt.Println("Time:", time.Since(start), "Part 1:", res)

	if part2 {
		start = time.Now()
		res := day.Solve2(inData, debug)
		fmt.Println("Time:", time.Since(start), "Part 2:", res)
	}
}

func getFileNames(day int) (input string, test string) {
	dayStr := strconv.Itoa(day)
	prefix := "inputs/day" + dayStr
	return prefix + "/in.txt", prefix + "/test.txt"
}

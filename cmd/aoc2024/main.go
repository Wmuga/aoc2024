package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	fileparser "github.com/wmuga/aoc2019/pkg/fileParser"
	"github.com/wmuga/aoc2024/internal/days"
	"github.com/wmuga/aoc2024/pkg/picker"
)

const (
	dayNum     = 3
	withPart2  = true
	toTest     = true
	debugInput = false
)

func getFileNames(day int) (input string, test string) {
	dayStr := strconv.Itoa(day)
	prefix := "inputs/day" + dayStr
	return prefix + "/in.txt", prefix + "/test.txt"
}

func main() {
	dayPicker := picker.NewDayRunner()
	days.Populate(dayPicker)
	// take correct day
	day, ok := dayPicker.GetDay(dayNum)
	if !ok {
		fmt.Printf("Day %d not found\n", dayNum)
		os.Exit(1)
	}

	// get input and test data
	inFile, testFile := getFileNames(dayNum)
	inData, err := fileparser.GetInput(inFile)
	if err != nil {
		fmt.Println("Error read input file:", err)
		os.Exit(1)
	}
	testData, err := fileparser.ReadTests(testFile)
	if err != nil {
		fmt.Println("Error read test file:", err)
		os.Exit(1)
	}

	doneTest := true
	// make tests
	if toTest {
		for _, test := range testData {
			var res string
			fmt.Println("Test", test.Name, "for part", test.Part)
			switch test.Part {
			case 1:
				res = day.Solve1(test.Data, true)
			case 2:
				if !withPart2 {
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

			out := "[OK]"
			if test.Answer != res {
				out = "[NO]"
				doneTest = false
			}

			fmt.Printf("%s expected: %s, result: %s\n", out, test.Answer, res)
		}
	}

	if !doneTest {
		os.Exit(1)
	}

	fmt.Println("\nAnswers:")

	start := time.Now()
	fmt.Println("Time:", time.Since(start), "Part 1:", day.Solve1(inData, debugInput))

	if withPart2 {
		start = time.Now()
		fmt.Println("Time:", time.Since(start), "Part 2:", day.Solve2(inData, debugInput))
	}
}

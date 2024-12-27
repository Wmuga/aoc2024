package days

import (
	"github.com/wmuga/aoc2024/pkg/models"
	"github.com/wmuga/aoc2024/pkg/picker"

	"github.com/wmuga/aoc2024/internal/days/day1"
	"github.com/wmuga/aoc2024/internal/days/day10"
	"github.com/wmuga/aoc2024/internal/days/day11"
	"github.com/wmuga/aoc2024/internal/days/day12"
	"github.com/wmuga/aoc2024/internal/days/day13"
	"github.com/wmuga/aoc2024/internal/days/day2"
	"github.com/wmuga/aoc2024/internal/days/day3"
	"github.com/wmuga/aoc2024/internal/days/day4"
	"github.com/wmuga/aoc2024/internal/days/day5"
	"github.com/wmuga/aoc2024/internal/days/day6"
	"github.com/wmuga/aoc2024/internal/days/day7"
	"github.com/wmuga/aoc2024/internal/days/day8"
	"github.com/wmuga/aoc2024/internal/days/day9"
)

var days = []models.Day{
	day1.Day{},
	day2.Day{},
	day3.Day{},
	day4.Day{},
	day5.Day{},
	day6.Day{},
	day7.Day{},
	day8.Day{},
	day9.Day{},
	day10.Day{},
	day11.Day{},
	day12.Day{},
	day13.Day{},
}

func Populate(d *picker.DayRunner) *picker.DayRunner {
	for i, day := range days {
		d.AddDay(picker.AoCDay{
			DayNum: i + 1,
			Solver: day,
		})
	}

	return d
}

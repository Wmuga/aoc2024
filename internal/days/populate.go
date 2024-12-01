package days

import (
	"github.com/wmuga/aoc2024/pkg/models"
	"github.com/wmuga/aoc2024/pkg/picker"
)

var days = []models.Day{}

func Populate(d *picker.DayRunner) *picker.DayRunner {
	for i, day := range days {
		d.AddDay(picker.AoCDay{
			DayNum: i + 1,
			Solver: day,
		})
	}

	return d
}

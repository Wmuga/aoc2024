package picker

import (
	"github.com/wmuga/aoc2024/pkg/models"
)

type DayRunner struct {
	days []models.Day
}

type AoCDay struct {
	DayNum int
	Solver models.Day
}

// Stud - day model for not implemented days
type Stud struct{}

// Solve1 implements models.Day.
func (s Stud) Solve1(input []string, debug bool) string {
	return ""
}

// Solve2 implements models.Day.
func (s Stud) Solve2(input []string, debug bool) string {
	return ""
}

var _ models.Day = Stud{}

func NewDayRunner(days ...AoCDay) *DayRunner {
	d := DayRunner{
		days: make([]models.Day, 25),
	}
	for _, day := range days {
		d.AddDay(day)
	}
	return &d
}

func (d *DayRunner) AddDay(day AoCDay) *DayRunner {
	d.days[day.DayNum-1] = day.Solver
	return d
}

func (d *DayRunner) GetDay(num int) (day models.Day, ok bool) {
	if num < 0 && num > 25 {
		return nil, false
	}

	day = d.days[num-1]
	if day == nil {
		return nil, false
	}

	// skip studs
	if _, ok := day.(Stud); ok {
		return nil, false
	}

	return day, true
}

func (d *DayRunner) CountDays() int {
	c := 0
	for _, day := range d.days {
		if day == nil {
			continue
		}
		c++
	}
	return c
}

package models

import "strconv"

type Day interface {
	Solve1(input []string, debug bool) string
	Solve2(input []string, debug bool) string
}

type Test struct {
	Name   string
	Part   int
	Answer string
	Data   []string
}

type Point2D struct {
	X int
	Y int
}

func (p Point2D) Add(o Point2D) Point2D {
	return Point2D{p.X + o.X, p.Y + o.Y}
}

func (p Point2D) Sub(o Point2D) Point2D {
	return Point2D{p.X - o.X, p.Y - o.Y}
}

func (p Point2D) InBoundOf(bound Point2D) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < bound.X && p.Y < bound.Y
}

func (p Point2D) String() string {
	return "{" + strconv.Itoa(p.X) + "; " + strconv.Itoa(p.Y) + "}"
}

type Point3D struct {
	X, Y, Z int
}

type PrintFunc func(string, ...interface{})

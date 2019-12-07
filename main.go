package main

import (
	"github.com/millera023/adventofcode/day1"
	"github.com/millera023/adventofcode/day2"
	"github.com/millera023/adventofcode/day3"
	"github.com/millera023/adventofcode/util"
)

func main() {
	day1.Run(util.Readlines("day1/input"))
	day2.Run(util.Readwords("day2/input"))
	day3.Run(util.Readlines("day3/input"))
}

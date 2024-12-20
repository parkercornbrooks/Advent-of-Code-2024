package main

import (
	"flag"
	"fmt"

	"github.com/parkercornbrooks/advent-of-code-2024/day1"
	"github.com/parkercornbrooks/advent-of-code-2024/day10"
	"github.com/parkercornbrooks/advent-of-code-2024/day11"
	"github.com/parkercornbrooks/advent-of-code-2024/day12"
	"github.com/parkercornbrooks/advent-of-code-2024/day13"
	"github.com/parkercornbrooks/advent-of-code-2024/day14"
	"github.com/parkercornbrooks/advent-of-code-2024/day15"
	"github.com/parkercornbrooks/advent-of-code-2024/day2"
	"github.com/parkercornbrooks/advent-of-code-2024/day3"
	"github.com/parkercornbrooks/advent-of-code-2024/day4"
	"github.com/parkercornbrooks/advent-of-code-2024/day5"
	"github.com/parkercornbrooks/advent-of-code-2024/day6"
	"github.com/parkercornbrooks/advent-of-code-2024/day7"
	"github.com/parkercornbrooks/advent-of-code-2024/day8"
	"github.com/parkercornbrooks/advent-of-code-2024/day9"
)

type dayPackage interface {
	Part1(day int, file string) int
	Part2(day int, file string) int
}

var packageMap = map[int]dayPackage{
	1:  day1.New(),
	2:  day2.New(),
	3:  day3.New(),
	4:  day4.New(),
	5:  day5.New(),
	6:  day6.New(),
	7:  day7.New(),
	8:  day8.New(),
	9:  day9.New(),
	10: day10.New(),
	11: day11.New(),
	12: day12.New(),
	13: day13.New(),
	14: day14.New(),
	15: day15.New(),
}

func main() {
	day := flag.Int("d", 1, "which day to run")
	part := flag.Int("p", 1, "which part of the given day to run")
	file := flag.String("f", "input.txt", "which input file to run")
	flag.Parse()

	res := run(*day, *part, *file)
	fmt.Println(res)
}

func run(day, part int, file string) int {
	d, ok := packageMap[day]
	if !ok {
		fmt.Printf("day %d does not exist\n", part)
		return 0
	}

	res := 0

	switch part {
	case 1:
		fmt.Printf("running day %d part %d with file %s\n", day, part, file)
		res = d.Part1(day, file)
	case 2:
		fmt.Printf("running day %d part %d with file %s\n", day, part, file)
		res = d.Part2(day, file)
	default:
		fmt.Printf("part %d does not exist\n", part)
	}
	return res
}

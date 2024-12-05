package day4

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

type Grid struct {
	m utils.Matrix
}

func (g Grid) scanFor(s string) int {
	starts := g.m.FindAll(s[:1])
	rest := s[1:]

	total := 0

	for _, start := range starts {
		total += g.checkSurrounding(start, rest)
	}

	return total
}

func (g Grid) checkSurrounding(start utils.Cell, rest string) int {
	found := 0
	surroundingCells := g.m.SurroundingCells(start)
	for _, step := range surroundingCells {
		match := g.checkDirection(step, rest)
		found += match
	}
	return found
}

func (g Grid) checkDirection(step utils.Step, rest string) int {
	currentStep := step
	reachedEdge := false
	for _, char := range strings.Split(rest, "") {
		if char != currentStep.Cell.Val || reachedEdge {
			return 0
		}
		nextCell, found := g.m.GetNext(currentStep.Cell, currentStep.Dir)
		if !found {
			reachedEdge = true
		}
		currentStep = utils.Step{
			Dir:  step.Dir,
			Cell: nextCell,
		}
	}
	return 1
}

func (d day) Part1(day int, file string) {
	grid := Grid{
		m: utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file)),
	}

	xmasFound := grid.scanFor("XMAS")
	fmt.Println(xmasFound)
}

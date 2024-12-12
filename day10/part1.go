package day10

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

func (d day) Part1(day int, file string) int {
	m := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	trailheads := m.FindAll("0")
	sum := 0

	for _, trailhead := range trailheads {
		pathEnds := map[utils.Cell]bool{
			trailhead: true,
		}
		for i := 0; i < 9; i++ {
			newCells := map[utils.Cell]bool{}
			for cell := range pathEnds {
				nextSteps := next(m, cell)
				for _, n := range nextSteps {
					newCells[n] = true
				}
			}
			pathEnds = newCells
		}
		sum += len(pathEnds)
	}

	return sum
}

func next(m utils.Matrix, c utils.Cell) []utils.Cell {
	nextInc := []utils.Cell{}
	next := m.SurroundingCells(c, "cardinal", false)
	for _, n := range next {
		if utils.MustAtoi(n.Cell.Val)-utils.MustAtoi(c.Val) == 1 {
			nextInc = append(nextInc, n.Cell)
		}
	}
	return nextInc
}

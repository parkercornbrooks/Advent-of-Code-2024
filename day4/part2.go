package day4

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (g Grid) scanForMas() int {
	found := 0
	starts := g.m.FindAll("A")
	for _, start := range starts {
		found += g.checkForX(start)
	}
	return found
}

func (g Grid) checkForX(start utils.Cell) int {
	diags := g.m.SurroundingCells(start, "diagonal")
	if len(diags) != 4 {
		return 0 // cell is on edge
	}
	sCount := 0
	mCount := 0
	for _, diag := range diags {
		if diag.Cell.Val == "M" {
			mCount++
		}
		if diag.Cell.Val == "S" {
			sCount++
		}
	}
	if sCount != 2 || mCount != 2 {
		return 0 // need 2 of each
	}
	if diags[0].Cell.Val == diags[3].Cell.Val {
		return 0 // M cannot be across from M, same with S
	}
	return 1
}

func (d day) Part2(day int, file string) int {
	grid := Grid{
		m: utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file)),
	}

	xmasFound := grid.scanForMas()
	return xmasFound
}

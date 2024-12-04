package main

import (
	"fmt"
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
	surrounding := g.m.SurroundingCells(start)
	if len(surrounding) != 8 {
		return 0 // cell is on edge
	}
	diags := []utils.Step{
		surrounding[0], // top left
		surrounding[2], // top right
		surrounding[5], // bottom left
		surrounding[7], // bottom right
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

func part2(day int, file string) {
	grid := Grid{
		m: utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file)),
	}

	xmasFound := grid.scanForMas()
	fmt.Println(xmasFound)
}

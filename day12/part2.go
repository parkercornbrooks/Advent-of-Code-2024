package day12

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) int {
	m := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	visited := map[utils.Cell]struct{}{}
	price := 0

	scanFunc := func(c utils.Cell) {
		if _, ok := visited[c]; ok {
			return
		}
		region, _, corners := flood(m, c, true)
		for cell := range region {
			visited[cell] = struct{}{}
		}
		price += corners * len(region)
	}
	m.ScanFunc(scanFunc)

	return price
}

func countCorners(m utils.Matrix, c utils.Cell) int {
	neighbors := 0
	corners := 0
	priorFound := false
	initialMatch := false
	checked := 0
	priorDiagNoMatch := false

	surr := m.SurroundingCells(c, "all", true)
	for i, step := range surr {
		if i%2 == 0 { // step is ordinal
			if step.Cell.Val == c.Val {
				if i == 0 {
					initialMatch = true
				}
				if priorFound {
					checked = 1
					if priorDiagNoMatch {
						corners++
					}
				}
				priorFound = true
				neighbors++
			} else {
				priorFound = false
			}
		} else { // step is diagonal
			priorDiagNoMatch = step.Cell.Val != c.Val
		}
	}

	if len(surr) == 8 && initialMatch && priorFound {
		checked = 1
		if priorDiagNoMatch {
			corners++
		}
	}

	switch neighbors {
	case 0:
		corners += 4
	case 1:
		corners += 2
	case 2:
		corners += checked
	}

	return corners
}

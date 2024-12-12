package day12

import (
	"fmt"
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
		region, _, corners := flood(m, c)
		for cell := range region {
			visited[cell] = struct{}{}
		}
		price += corners * len(region)
		fmt.Printf("Region %s: %d * %d = %d\n", c.Val, len(region), corners, corners*len(region))
	}
	m.ScanFunc(scanFunc)

	return price
}

func countCorners(m utils.Matrix, c utils.Cell, surr []utils.Step) int {
	diags := m.SurroundingCells(c, "diagonal")
	neighbors := 0
	corners := 0
	priorFound := false
	initialMatch := false
	checked := 0
	for i := 0; i < len(surr); i++ {
		if c.Val == surr[i].Cell.Val {
			neighbors++
			if priorFound {
				checked++
				if diags[i-1].Cell.Val != c.Val {
					corners++
				}
			}
			if i == 0 {
				initialMatch = true
			}
			priorFound = true
		} else {
			priorFound = false
		}
	}
	if len(surr) == 4 && priorFound && initialMatch {
		checked++
		if diags[3].Cell.Val != c.Val {
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
	if c.Val == "M" {
		fmt.Printf("M at %s has %d corners\n", c, corners)
	}
	return corners
}

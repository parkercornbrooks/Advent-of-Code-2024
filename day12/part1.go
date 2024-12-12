package day12

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

	visited := map[utils.Cell]struct{}{}
	price := 0

	scanFunc := func(c utils.Cell) {
		if _, ok := visited[c]; ok {
			return
		}
		region, edges := flood(m, c)
		for cell := range region {
			visited[cell] = struct{}{}
		}
		price += edges * len(region)
	}
	m.ScanFunc(scanFunc)

	return price
}
func flood(m utils.Matrix, c utils.Cell) (map[utils.Cell]struct{}, int) {
	visited := map[utils.Cell]struct{}{c: {}}
	latest := map[utils.Cell]struct{}{c: {}}
	edges := 0

	for {
		newVisits := map[utils.Cell]struct{}{}
		for cell := range latest {
			edges += 4
			surr := m.SurroundingCells(cell, "cardinal")
			for _, s := range surr {
				if s.Cell.Val == c.Val {
					edges -= 1
					if _, ok := visited[s.Cell]; !ok {
						visited[s.Cell] = struct{}{}
						newVisits[s.Cell] = struct{}{}
					}
				}
			}
		}
		if len(newVisits) == 0 {
			break
		}
		latest = newVisits
	}

	return visited, edges
}

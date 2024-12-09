package day8

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func getAntinodesForPair(g utils.Matrix, a, b utils.Cell) []string {
	rDiff := a.R - b.R
	cDiff := a.C - b.C
	r, c := a.R, a.C

	antinodes := []string{a.String(), b.String()}

	for {
		newR, newC := r+rDiff, c+cDiff
		cell, exists := g.GetCell(newR, newC)
		if !exists {
			break
		}
		antinodes = append(antinodes, cell.String())
		r, c = newR, newC
	}
	r, c = b.R, b.C
	for {
		newR, newC := r-rDiff, c-cDiff
		cell, exists := g.GetCell(newR, newC)
		if !exists {
			break
		}
		antinodes = append(antinodes, cell.String())
		r, c = newR, newC
	}

	return antinodes
}

func (d day) Part2(day int, file string) {
	grid := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	antennae := buildAntennaMap(grid)
	antinodes := make(map[string]bool)

	for _, cells := range antennae {
		for i := 0; i < len(cells)-1; i++ {
			for _, b := range cells[i+1:] {
				for _, n := range getAntinodesForPair(grid, cells[i], b) {
					antinodes[n] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))

}

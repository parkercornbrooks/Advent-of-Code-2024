package day8

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

func reflect(m utils.Matrix, a, b utils.Cell) (utils.Cell, bool) {
	rDiff := a.R - b.R
	newR := a.R + rDiff
	cDiff := a.C - b.C
	newC := a.C + cDiff
	if _, exists := m.GetCell(newR, newC); !exists {
		return a, false
	}
	return m[newR][newC], true
}

func (d day) Part1(day int, file string) int {
	grid := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	antennae := buildAntennaMap(grid)
	antinodes := make(map[string]bool)

	for _, cells := range antennae {
		for i := 0; i < len(cells)-1; i++ {
			for _, b := range cells[i+1:] {
				c, exists := reflect(grid, cells[i], b)
				if exists {
					antinodes[c.String()] = true
				}
				c, exists = reflect(grid, b, cells[i])
				if exists {
					antinodes[c.String()] = true
				}
			}
		}
	}

	return len(antinodes)

}

func buildAntennaMap(g utils.Matrix) map[string][]utils.Cell {
	antennae := map[string][]utils.Cell{}
	g.ScanFunc(func(c utils.Cell) {
		if c.Val != "." {
			antennae[c.Val] = append(antennae[c.Val], c)
		}
	})
	return antennae
}

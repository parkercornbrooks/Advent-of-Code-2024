package day8

import (
	"fmt"
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
	if newR < 0 || newR >= m.Height() {
		return a, false
	}
	cDiff := a.C - b.C
	newC := a.C + cDiff
	if newC < 0 || newC >= m.Width() {
		return a, false
	}
	return m[newR][newC], true
}

func (d day) Part1(day int, file string) {
	grid := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	antennae := map[string][]utils.Cell{}
	grid.ScanFunc(func(c utils.Cell) {
		if c.Val != "." {
			antennae[c.Val] = append(antennae[c.Val], c)
		}
	})

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

	fmt.Println(len(antinodes))

}

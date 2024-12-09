package day6

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) int {
	grid := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	loc := grid.FindAll("^")[0] // my guard is a caret so skip check for ><V
	visited := map[string]bool{
		loc.String(): true,
	}
	dir := utils.DirMap["up"]
	loops := 0

	for {
		nextCell, exists := grid.GetNext(loc, dir)
		if !exists {
			break
		}
		if nextCell.Val == "#" {
			dir = turnRight(dir)
		} else {
			if _, ok := visited[nextCell.String()]; !ok {
				visited[nextCell.String()] = true
				loops += scenario(grid, loc, nextCell, dir)
			}
			loc = nextCell
		}
	}

	return loops
}

// scenario returns 1 if a loop is found, else 0
func scenario(m utils.Matrix, loc, next utils.Cell, dir utils.Dir) int {
	m = m.Copy()
	m[next.R][next.C].Val = "#"

	tag := makeTag(loc, dir)

	visited := map[string]bool{
		tag: true,
	}

	for {
		nextCell, exists := m.GetNext(loc, dir)
		if !exists {
			return 0 // reached the end of the grid and are not in a loop
		}
		if nextCell.Val == "#" {
			dir = turnRight(dir)
		} else {
			tag = makeTag(loc, dir)
			if _, ok := visited[tag]; ok {
				return 1 // visited this cell in the same direction, therefore are looping
			}
			visited[tag] = true
			loc = nextCell
		}
	}
}

// makeTag makes a unique tag with a cell and given direction
func makeTag(loc utils.Cell, dir utils.Dir) string {
	return fmt.Sprintf("%s:%d,%d", loc, dir.R, dir.C)
}

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
	dir := utils.DirMap["u"]
	loops := 0

	results := make(chan int)
	sem := make(chan struct{}, 8)

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
				go func(m utils.Matrix, loc, next utils.Cell, dir utils.Dir) {
					sem <- struct{}{}
					results <- scenario(m, loc, next, dir)
					<-sem
				}(grid, loc, nextCell, dir)
			}
			loc = nextCell
		}
	}

	for i := 0; i < len(visited)-1; i++ {
		res := <-results
		loops += res
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

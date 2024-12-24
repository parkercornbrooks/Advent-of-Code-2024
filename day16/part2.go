package day16

import (
	"math"
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type Path struct {
	head  utils.Cell
	cells []utils.Cell
	dir   utils.Dir
	score int
}

func (d day) Part2(day int, file string) int {
	m := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))
	startDir := utils.DirMap["r"]
	start := m.FindAll(START)[0]
	move1 := Path{
		head:  start,
		cells: []utils.Cell{start},
		dir:   startDir,
		score: 0,
	}
	paths := []Path{move1}
	shortest := math.MaxInt
	found := map[utils.Cell]int{}
	allShortest := []Path{}
	for {
		newPaths := []Path{}
		for _, path := range paths {
			if path.head.Val == END && path.score <= shortest {
				allShortest = append(allShortest, path)
				shortest = path.score
			} else {
				if ct, exists := found[path.head]; exists && ct+1001 < path.score {
					continue
				}
				found[path.head] = path.score
				if path.score < shortest {
					nextPaths := advance(m, path)
					newPaths = append(newPaths, nextPaths...)
				}
			}
		}
		if len(newPaths) == 0 {
			break
		}
		paths = newPaths
	}

	seats := make(map[utils.Cell]struct{})

	for _, p := range allShortest {
		if p.score != shortest {
			continue
		}
		for _, c := range p.cells {
			seats[c] = struct{}{}
		}
	}

	return len(seats)
}

func advance(m utils.Matrix, current Path) []Path {
	paths := []Path{}
	for _, direction := range utils.Cardinals {
		if direction == current.dir.Opposite() {
			continue
		}
		next, _ := m.GetNext(current.head, direction)
		if next.Val == BLANK || next.Val == END {
			dirAddition := 1000
			if current.dir == direction {
				dirAddition = 0
			}
			cells := copyAndAdd(current.cells, next)
			path := Path{
				head:  next,
				cells: cells,
				dir:   direction,
				score: current.score + 1 + dirAddition,
			}
			paths = append(paths, path)
		}
	}
	return paths
}

func copyAndAdd(cells []utils.Cell, cell utils.Cell) []utils.Cell {
	newC := make([]utils.Cell, len(cells)+1)
	copy(newC, cells)
	newC[len(cells)] = cell
	return newC
}

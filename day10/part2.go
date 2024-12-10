package day10

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) int {
	m := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	trailheads := m.FindAll("0")
	sum := 0

	for _, trailhead := range trailheads {
		paths := [][]utils.Cell{{trailhead}}
		for i := 0; i < 9; i++ {
			newPaths := [][]utils.Cell{}
			for _, path := range paths {
				nextSteps := next(m, path[i])
				for _, n := range nextSteps {
					nextPath := make([]utils.Cell, i+2)
					copy(nextPath, path)
					nextPath[i+1] = n
					newPaths = append(newPaths, nextPath)
				}
			}
			paths = newPaths
		}
		sum += len(paths)
	}

	return sum
}

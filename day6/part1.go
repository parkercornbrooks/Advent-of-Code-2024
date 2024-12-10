package day6

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
	grid := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	loc := grid.FindAll("^")[0] // my guard is a caret so skip check for ><V
	visited := map[string]bool{
		loc.String(): true,
	}
	dir := utils.DirMap["u"]

	for {
		nextCell, exists := grid.GetNext(loc, dir)
		if !exists {
			break
		}
		if nextCell.Val == "#" {
			dir = turnRight(dir)
		} else {
			loc = nextCell
			if _, ok := visited[loc.String()]; !ok {
				visited[loc.String()] = true
			}
		}
	}

	return len(visited)
}

func turnRight(dir utils.Dir) utils.Dir {
	switch dir {
	case utils.DirMap["u"]:
		return utils.DirMap["r"]
	case utils.DirMap["r"]:
		return utils.DirMap["d"]
	case utils.DirMap["d"]:
		return utils.DirMap["l"]
	case utils.DirMap["l"]:
		return utils.DirMap["u"]
	default: // should not hit this case
		return dir
	}
}

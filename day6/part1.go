package day6

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

func (d day) Part1(day int, file string) {
	grid := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	loc := grid.FindAll("^")[0] // my guard is a caret so skip check for ><V
	visited := map[string]bool{
		loc.String(): true,
	}
	dir := utils.DirMap["up"]

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

	fmt.Println(len(visited))
}

func turnRight(dir utils.Dir) utils.Dir {
	switch dir {
	case utils.DirMap["up"]:
		return utils.DirMap["right"]
	case utils.DirMap["right"]:
		return utils.DirMap["down"]
	case utils.DirMap["down"]:
		return utils.DirMap["left"]
	case utils.DirMap["left"]:
		return utils.DirMap["up"]
	default: // should not hit this case
		return dir
	}
}

package day15

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

const EDGE = "#"
const BOX = "O"
const EMPTY = "."
const BOT = "@"

var dirMap = map[string]utils.Dir{
	"^": utils.DirMap["u"],
	">": utils.DirMap["r"],
	"v": utils.DirMap["d"],
	"<": utils.DirMap["l"],
}

type warehouse struct {
	grid utils.Matrix
	bot  utils.Cell
}

func (w *warehouse) step(arrow string) {
	dir := dirMap[arrow]
	current := w.bot
	cells := []utils.Cell{current}
	for {
		current, _ = w.grid.GetNext(current, dir)
		if current.Val == EDGE {
			break
		}
		if current.Val == EMPTY {
			w.move(cells, dir)
			break
		}
		cells = append(cells, current)
	}
}

func (w *warehouse) gpsSum() int {
	total := 0
	w.grid.ScanFunc(func(c utils.Cell) {
		if c.Val == BOX || c.Val == BOXL {
			coord := c.R*100 + c.C
			total += coord
		}
	})

	return total
}

func (d day) Part1(day int, file string) int {
	m, ins := load(day, file, false)
	w := warehouse{grid: m, bot: m.FindAll(BOT)[0]}
	for i := range ins {
		w.step(ins[i : i+1])
	}

	return w.gpsSum()
}

func load(day int, file string, expanded bool) (utils.Matrix, string) {
	isGrid := true
	m := make(utils.Matrix, 0)
	rowNum := 0
	instructions := ""

	linefn := func(line string) {
		if line == "" {
			isGrid = false
		} else if isGrid {
			row := []utils.Cell{}
			for i, char := range line {
				if expanded {
					newCells := expandedCellMap[string(char)]
					cell1 := utils.Cell{R: rowNum, C: i * 2, Val: newCells[0]}
					cell2 := utils.Cell{R: rowNum, C: i*2 + 1, Val: newCells[1]}
					row = append(row, cell1, cell2)
				} else {
					cell := utils.Cell{R: rowNum, C: i, Val: string(char)}
					row = append(row, cell)
				}
			}
			m = append(m, row)
			rowNum++
		} else {
			instructions += line
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return m, instructions
}

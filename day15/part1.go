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
		cells = append(cells, current)
		if current.Val == EMPTY {
			for i, c := range cells {
				switch i {
				case 0:
					w.grid[c.R][c.C].Val = EMPTY
				case 1:
					w.grid[c.R][c.C].Val = BOT
					w.bot = c
				default:
					w.grid[c.R][c.C].Val = BOX
				}
			}
			break
		}
	}
}

func (d day) Part1(day int, file string) int {
	m, ins := load(day, file)
	w := warehouse{grid: m, bot: m.FindAll(BOT)[0]}
	for i := range ins {
		w.step(ins[i : i+1])
	}

	total := 0
	w.grid.ScanFunc(func(c utils.Cell) {
		if c.Val == BOX {
			coord := c.R*100 + c.C
			total += coord
		}
	})

	return total
}

func load(day int, file string) (utils.Matrix, string) {
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
				cell := utils.Cell{R: rowNum, C: i, Val: string(char)}
				row = append(row, cell)
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

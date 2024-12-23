package day15

import (
	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

const BOXL = "["
const BOXR = "]"

var expandedCellMap = map[string][]string{
	EDGE:  {EDGE, EDGE},
	BOX:   {BOXL, BOXR},
	EMPTY: {EMPTY, EMPTY},
	BOT:   {BOT, EMPTY},
}

func (w *warehouse) move(cells []utils.Cell, dir utils.Dir) {
	for i := len(cells); i > 0; i-- {
		c := cells[i-1]
		w.grid[c.R+dir.R][c.C+dir.C].Val = c.Val
		w.grid[c.R][c.C].Val = EMPTY
	}
	w.bot, _ = w.grid.GetNext(cells[0], dir)
}

func (w *warehouse) step2(arrow string) {
	if arrow == ">" || arrow == "<" {
		w.step(arrow)
		return
	}
	dir := dirMap[arrow]
	iter := []utils.Cell{w.bot}
	all := []utils.Cell{w.bot}
	for {
		nextIter := make([]utils.Cell, 0)
		for _, cell := range iter {
			next, _ := w.grid.GetNext(cell, dir)
			if next.Val == EDGE {
				return
			} else if next.Val == BOXR {
				other, _ := w.grid.GetNext(next, utils.DirMap["l"])
				nextIter = append(nextIter, next, other)
			} else if next.Val == BOXL {
				other, _ := w.grid.GetNext(next, utils.DirMap["r"])
				nextIter = append(nextIter, next, other)
			}
		}
		if len(nextIter) == 0 {
			w.move(all, dir)
			return
		}
		all = append(all, nextIter...)
		iter = nextIter
	}
}

func (d day) Part2(day int, file string) int {
	m, ins := load(day, file, true)
	w := warehouse{grid: m, bot: m.FindAll(BOT)[0]}

	for i := range ins {
		w.step2(ins[i : i+1])
	}

	return w.gpsSum()
}

package day14

import (
	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (r *Robot) move(x, y int) {
	r.x = x
	r.y = y
}

func (a *Area) step() {
	m := newGrid(a.w, a.h)
	for _, r := range a.bots {
		r.move(a.reposition(r.posAfterSec(1)))
		m[r.y][r.x].Val = "X"
	}
	for _, r := range m {
		ct := 0
		for _, c := range r {
			if c.Val == "X" {
				ct++
			} else {
				ct = 0
			}
			if ct > 10 {
				// There are 10 bots next to each other in a row
				a.rowMaxReached = true
			}
		}
	}
	// if a.rowMaxReached {
	// 	fmt.Print(m)
	// }
}

func newGrid(w, h int) utils.Matrix {
	m := make(utils.Matrix, h)
	for r := range h {
		m[r] = make([]utils.Cell, w)
		for c := range w {
			m[r][c] = utils.Cell{R: r, C: c, Val: "."}
		}
	}
	return m
}

func (d day) Part2(day int, file string) int {
	if file == "example.txt" {
		return 0
	}

	robots := load(day, file)
	a := &Area{101, 103, robots, false}

	steps := 0

	for {
		a.step()
		steps++
		if a.rowMaxReached {
			break
		}
	}

	return steps
}

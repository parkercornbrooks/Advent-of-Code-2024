package utils

import "fmt"

type Cell struct {
	R, C int
	Val  string
}

func (c Cell) String() string {
	return fmt.Sprintf("%d-%d", c.R, c.C)
}

type Matrix [][]Cell

type Dir struct {
	R, C int
}

var DirMap = map[string]Dir{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},
}

var directions = []Dir{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Step struct {
	Dir  Dir
	Cell Cell
}

func (m Matrix) width() int {
	return len(m[0])
}

func (m Matrix) height() int {
	return len(m)
}

func (m Matrix) Transposed() Matrix {
	newMatrix := make(Matrix, 0)
	for i := 0; i < m.width(); i++ {
		row := make([]Cell, 0)
		for j := 0; j < m.height(); j++ {
			row = append(row, m[j][i])
		}
		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

func (m Matrix) FindAll(s string) []Cell {
	found := []Cell{}
	for r := range m.height() {
		for c := range m.width() {
			cell := m[r][c]
			if cell.Val == s {
				found = append(found, cell)
			}
		}
	}
	return found
}

func (m Matrix) SurroundingCells(c Cell) []Step {
	steps := []Step{}
	for _, dir := range directions {
		cell, exists := m.GetNext(c, dir)
		if exists {
			steps = append(steps, Step{
				Dir:  dir,
				Cell: cell,
			})
		}
	}
	return steps
}

func (m Matrix) GetCell(r, c int) (Cell, bool) {
	if r < 0 || r == m.height() || c < 0 || c == m.width() {
		return Cell{}, false
	}
	return m[r][c], true
}

func (m Matrix) GetNext(c Cell, d Dir) (Cell, bool) {
	newR := c.R + d.R
	newC := c.C + d.C
	return m.GetCell(newR, newC)
}

func (m Matrix) Copy() Matrix {
	dupe := make([][]Cell, len(m))
	for i := range m {
		dupe[i] = make([]Cell, len(m[i]))
		copy(dupe[i], m[i])
	}
	return dupe
}

func (m Matrix) String() string {
	s := ""
	for _, row := range m {
		for _, cell := range row {
			s += cell.Val
		}
		s += "\n"
	}
	return s
}

func ReadIntoMatrix(filepath string) Matrix {
	matrix := make(Matrix, 0)
	rowNum := 0

	linefn := func(line string) {
		row := []Cell{}
		for i, char := range line {
			cell := Cell{rowNum, i, string(char)}
			row = append(row, cell)
		}
		matrix = append(matrix, row)
		rowNum++
	}

	ReadInput(filepath, linefn)

	return matrix
}

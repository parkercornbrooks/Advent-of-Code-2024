package utils

import "fmt"

const EMPTY_CELL_VAL = "empty"

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

func (d Dir) Opposite() Dir {
	return Dir{R: -d.R, C: -d.C}
}

var DirMap = map[string]Dir{
	"u-l": {-1, -1},
	"u":   {-1, 0},
	"u-r": {-1, 1},
	"l":   {0, -1},
	"r":   {0, 1},
	"d-l": {1, -1},
	"d":   {1, 0},
	"d-r": {1, 1},
}

var Cardinals = []Dir{DirMap["u"], DirMap["r"], DirMap["d"], DirMap["l"]}
var Diagonals = []Dir{DirMap["u-r"], DirMap["d-r"], DirMap["d-l"], DirMap["u-l"]}
var AllDirs = []Dir{DirMap["u"], DirMap["u-r"], DirMap["r"], DirMap["d-r"],
	DirMap["d"], DirMap["d-l"], DirMap["l"], DirMap["u-l"]}

type Step struct {
	Dir  Dir
	Cell Cell
}

func (m Matrix) Width() int {
	return len(m[0])
}

func (m Matrix) Height() int {
	return len(m)
}

func (m Matrix) Transposed() Matrix {
	newMatrix := make(Matrix, 0)
	for i := 0; i < m.Width(); i++ {
		row := make([]Cell, 0)
		for j := 0; j < m.Height(); j++ {
			row = append(row, m[j][i])
		}
		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

func (m Matrix) ScanFunc(f func(Cell)) {
	for r := range m.Height() {
		for c := range m.Width() {
			f(m[r][c])
		}
	}
}

func (m Matrix) FindAll(s string) []Cell {
	found := []Cell{}
	m.ScanFunc(func(c Cell) {
		if c.Val == s {
			found = append(found, c)
		}
	})
	return found
}

func (m Matrix) SurroundingCells(c Cell, which string, fill bool) []Step {
	var dirs []Dir
	switch which {
	case "diagonal":
		dirs = Diagonals
	case "cardinal":
		dirs = Cardinals
	default:
		dirs = AllDirs
	}
	return m.getNearby(c, dirs, fill)
}

func (m Matrix) getNearby(c Cell, dirs []Dir, fill bool) []Step {
	steps := []Step{}
	for _, dir := range dirs {
		cell, exists := m.GetNext(c, dir)
		if exists {
			steps = append(steps, Step{
				Dir:  dir,
				Cell: cell,
			})
		} else if fill {
			steps = append(steps, Step{
				Dir: dir,
				Cell: Cell{
					Val: EMPTY_CELL_VAL,
				},
			})
		}
	}
	return steps
}

func (m Matrix) GetCell(r, c int) (Cell, bool) {
	if r < 0 || r >= m.Height() || c < 0 || c >= m.Width() {
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

package day16

import (
	"math"
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

const WALL = "#"
const BLANK = "."
const START = "S"
const END = "E"

type Move struct {
	cell  utils.Cell
	dir   utils.Dir
	score int
}

func (d day) Part1(day int, file string) int {
	m := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))
	startDir := utils.DirMap["r"]
	start := m.FindAll(START)[0]
	move1 := Move{cell: start, dir: startDir, score: 0}
	heads := []Move{move1}
	shortest := math.MaxInt
	found := map[utils.Cell]int{}
	for {
		newHeads := []Move{}
		for _, head := range heads {
			if head.cell.Val == END && head.score < shortest {
				shortest = head.score
			} else {
				if ct, exists := found[head.cell]; exists && ct < head.score {
					continue
				}
				found[head.cell] = head.score
				if head.score < shortest {
					nextMoves := step(m, head)
					newHeads = append(newHeads, nextMoves...)
				}
			}
		}
		if len(newHeads) == 0 {
			break
		}
		heads = newHeads
	}

	return shortest
}

func step(m utils.Matrix, current Move) []Move {
	moves := []Move{}
	for _, direction := range utils.Cardinals {
		if direction == current.dir.Opposite() {
			continue
		}
		next, _ := m.GetNext(current.cell, direction)
		if next.Val == BLANK || next.Val == END {
			dirAddition := 1000
			if current.dir == direction {
				dirAddition = 0
			}
			move := Move{
				cell:  next,
				dir:   direction,
				score: current.score + 1 + dirAddition,
			}
			moves = append(moves, move)
		}
	}
	return moves
}

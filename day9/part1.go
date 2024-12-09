package day9

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

type block struct {
	fileId int
	empty  bool
}

func newEmptyBlock() block {
	return block{fileId: 0, empty: true}
}

func newBlock(id int) block {
	return block{fileId: id, empty: false}
}

func (d day) Part1(day int, file string) int {
	disk := make([]block, 0)
	fileBlockCt := 0

	lineFunc := func(l string) {
		fileId := 0
		for i, v := range strings.Split(l, "") {
			length := utils.MustAtoi(v)
			toAppend := make([]block, length)

			if i%2 == 0 { // this is a file
				for p := 0; p < length; p++ {
					toAppend[p] = newBlock(fileId)
				}
				fileId++
				fileBlockCt += length
			} else { // this is a space
				for p := 0; p < length; p++ {
					toAppend[p] = newEmptyBlock()
				}
			}
			disk = append(disk, toAppend...)
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), lineFunc)

	total := 0
	popInd := len(disk) - 1
	for i := range fileBlockCt {
		current := disk[i]
		val := current.fileId
		if disk[i].empty {
			for pi := popInd; pi > 0; pi-- {
				if !disk[pi].empty {
					val = disk[pi].fileId
					popInd = pi - 1
					break
				}
			}
		}
		total += i * val
	}

	return total
}

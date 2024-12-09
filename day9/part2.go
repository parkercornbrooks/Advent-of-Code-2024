package day9

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) int {
	files, spaces := []int{}, []int{}
	linefn := func(line string) {
		for i, v := range strings.Split(line, "") {
			if i%2 == 0 {
				files = append(files, utils.MustAtoi(v))
			} else {
				spaces = append(spaces, utils.MustAtoi(v))
			}
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)
	fileMoves := make(map[int][]int)

	for i := len(files) - 1; i > 0; i-- {
		file := files[i]
		for j := 0; j < i; j++ {
			if file <= spaces[j] {
				spaces[j] = spaces[j] - file
				fileMoves[j] = append(fileMoves[j], i)
				break
			}
		}
	}

	disk := span(files[0], 0)
	moveTracker := make([]bool, len(files))

	for i := 0; i < len(files)-1; i++ {
		movedFiles, exists := fileMoves[i]
		if exists {
			for _, fIndex := range movedFiles {
				// add moved files
				disk = append(disk, span(files[fIndex], fIndex)...)
				moveTracker[fIndex] = true
			}
		}
		// add remaining spaces
		disk = append(disk, span(spaces[i], 0)...)
		//add in-place files
		val := i + 1
		if moveTracker[val] {
			val = 0
		}
		disk = append(disk, span(files[i+1], val)...)
	}

	total := 0

	for i := 0; i < len(disk); i++ {
		total += i * disk[i]
	}

	return total
}

func span(len, val int) []int {
	span := make([]int, len)
	for i := range len {
		span[i] = val
	}
	return span
}

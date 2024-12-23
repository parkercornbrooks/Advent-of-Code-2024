package day16

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) int {
	linefn := func(line string) {}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return 0
}

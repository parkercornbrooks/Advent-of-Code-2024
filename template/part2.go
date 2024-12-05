package main

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func part2(day int, file string) {
	linefn := func(line string) {

	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

}

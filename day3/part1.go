package day3

import (
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

var r = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func (d day) Part1(day int, file string) int {
	total := 0

	linefn := func(line string) {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			v1, v2 := utils.MustAtoi(m[1]), utils.MustAtoi(m[2])
			product := v1 * v2
			total += product
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return total
}

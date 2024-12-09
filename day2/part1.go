package day2

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

type report []int

func (r report) isSafe() bool {
	inc := true
	for i := 0; i < len(r)-1; i++ {
		diff := r[i+1] - r[i]
		isInc := diff < 0
		if i == 0 {
			inc = isInc
		}
		if isInc != inc {
			return false
		}
		abs := utils.Abs(diff)
		if abs < 1 || abs > 3 {
			return false
		}
	}
	return true
}

func createReport(raw []string) report {
	report := make([]int, len(raw))
	for ind, s := range raw {
		num := utils.MustAtoi(s)
		report[ind] = num
	}
	return report
}

func (d day) Part1(day int, file string) int {
	safeReports := 0

	linefn := func(line string) {
		levels := strings.Fields(line)
		r := createReport(levels)
		if r.isSafe() {
			safeReports += 1
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return safeReports
}

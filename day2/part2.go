package day2

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (r report) isSafeWithDampener() bool {
	if r.isSafe() {
		return true
	}
	for i := range len(r) {
		newReport := newReportWithoutIndex(r, i)
		if newReport.isSafe() {
			return true
		}
	}
	return false
}

func (d day) Part2(day int, file string) int {
	safeReports := 0

	linefn := func(line string) {
		levels := strings.Fields(line)
		r := createReport(levels)
		if r.isSafeWithDampener() {
			safeReports += 1
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return safeReports
}

func newReportWithoutIndex(r report, index int) report {
	var nr []int
	for i := range len(r) {
		if i != index {
			nr = append(nr, r[i])
		}
	}
	return report(nr)
}

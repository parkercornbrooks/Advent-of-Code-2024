package main

import (
	"fmt"
	"path/filepath"
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

func part2() {
	fmt.Println("day 2 part 2")

	safeReports := 0

	linefn := func(line string) {
		levels := strings.Fields(line)
		r := createReport(levels)
		if r.isSafeWithDampener() {
			safeReports += 1
		}
	}

	endfn := func() {
		fmt.Printf("Total safe reports: %d\n", safeReports)
	}

	utils.ReadInput(filepath.Join("day2", "input.txt"), linefn, endfn)
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
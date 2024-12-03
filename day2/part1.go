package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

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
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		report[ind] = num
	}
	return report
}

func part1(day int, file string) {
	safeReports := 0

	linefn := func(line string) {
		levels := strings.Fields(line)
		r := createReport(levels)
		if r.isSafe() {
			safeReports += 1
		}
	}

	endfn := func() {
		fmt.Printf("Total safe reports: %d\n", safeReports)
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn, endfn)
}

package day1

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

func (d day) Part1(day int, file string) {
	left := make([]int, 1000)
	right := make([]int, 1000)
	diff := 0

	linefn := func(line string) {
		entries := strings.Fields(line)
		l, r := entries[0], entries[1]
		lVal, rVal := utils.MustAtoi(l), utils.MustAtoi(r)
		left = append(left, lVal)
		right = append(right, rVal)
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	slices.Sort(left)
	slices.Sort(right)
	for i := 0; i < len(left); i++ {
		diff += utils.Abs(left[i] - right[i])
	}
	fmt.Printf("Total: %d\n", diff)
}

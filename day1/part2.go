package day1

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) {
	left := make([]int, 1000)
	right := make(map[int]int)

	linefn := func(line string) {
		entries := strings.Fields(line)
		l, r := entries[0], entries[1]
		lVal, rVal := utils.MustAtoi(l), utils.MustAtoi(r)
		left = append(left, lVal)
		right[rVal] = right[rVal] + 1
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	sim := 0

	for _, num := range left {
		sim += num * right[num]
	}
	fmt.Printf("Total: %d\n", sim)
}

package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func part1() {
	fmt.Println("day 1 part 1")

	left := make([]int, 1000)
	right := make([]int, 1000)
	diff := 0

	linefn := func(line string) {
		entries := strings.Fields(line)
		l, r := entries[0], entries[1]
		lVal, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		rVal, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		left = append(left, lVal)
		right = append(right, rVal)
	}

	endfn := func() {
		slices.Sort(left)
		slices.Sort(right)
		for i := 0; i < len(left); i++ {
			diff += abs(left[i] - right[i])
		}
		fmt.Printf("Total: %d\n", diff)
	}

	utils.ReadInput(filepath.Join("day1", "input.txt"), linefn, endfn)
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}

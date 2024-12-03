package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func part2() {
	fmt.Println("day 1 part 2")

	left := make([]int, 1000)
	right := make(map[int]int)

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
		right[rVal] = right[rVal] + 1
	}

	endfn := func() {
		sim := 0

		for _, num := range left {
			sim += num * right[num]
		}
		fmt.Printf("Total: %d\n", sim)
	}

	utils.ReadFile(filepath.Join("day1", "input.txt"), linefn, endfn)
}

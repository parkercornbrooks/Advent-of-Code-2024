package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

var r2 = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)

func part2(day int, file string) {
	total := 0
	use := true

	linefn := func(line string) {
		matches := r2.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			switch m[0][:3] {
			case "don":
				use = false
			case "do(":
				use = true
			case "mul":
				if use {
					v1, v2 := utils.MustAtoi(m[2]), utils.MustAtoi(m[3])
					product := v1 * v2
					total += product
				}
			}
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	fmt.Printf("Total: %d\n", total)
}

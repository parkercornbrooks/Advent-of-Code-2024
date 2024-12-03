package main

import (
	"flag"
	"fmt"
)

var day = 3

func main() {
	part := flag.Int("p", 1, "which part of the given day to run")
	file := flag.String("f", "input.txt", "which input file to run")
	flag.Parse()

	switch *part {
	case 1:
		fmt.Printf("running day %d part %d with file %s\n", day, *part, *file)
		part1(day, *file)
	case 2:
		fmt.Printf("running day %d part %d with file %s\n", day, *part, *file)
		part2(day, *file)
	default:
		fmt.Printf("part %d does not exist\n", *part)
	}
}

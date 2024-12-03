package main

import (
	"flag"
	"fmt"
)

func main() {
	part := flag.Int("part", 1, "which part of the given day to run")
	flag.Parse()

	switch *part {
	case 1:
		part1()
	default:
		fmt.Printf("part %d does not exist", *part)
	}
}

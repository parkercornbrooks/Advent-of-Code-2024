package day7

import "fmt"

func (d day) Part2(day int, file string) {
	eqs := readInput(day, file)
	total := 0
	ops := []string{Star, Plus, Conc}

	for _, e := range eqs {
		total += e.isValid(ops)
	}

	fmt.Println(total)
}

package day7

func (d day) Part2(day int, file string) int {
	eqs := readInput(day, file)
	total := 0
	ops := []string{Star, Plus, Conc}

	for _, e := range eqs {
		total += e.isValid(ops)
	}

	return total
}

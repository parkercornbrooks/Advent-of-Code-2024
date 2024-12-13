package day13

const POS_CORR = 10_000_000_000_000

func (d day) Part2(day int, file string) int {
	machines := load(day, file)
	total := 0

	for _, m := range machines {
		m.px += POS_CORR
		m.py += POS_CORR
		a, b := m.play()

		total += a*A_SCORE + b*B_SCORE
	}
	return total
}

package day5

func (d day) Part2(day int, file string) int {
	rules, updates := loadData(day, file)

	total := 0

	for _, update := range updates {
		if !rules.orderIsValid(update) {
			newOrder := rules.arrange(update)
			total += middlePage(newOrder)
		}
	}
	return total
}

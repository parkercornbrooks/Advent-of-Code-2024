package main

import (
	"fmt"
)

func part2(day int, file string) {
	rules, updates := loadData(day, file)

	total := 0

	for _, update := range updates {
		if !rules.orderIsValid(update) {
			newOrder := rules.arrange(update)
			total += middlePage(newOrder)
		}
	}
	fmt.Println("Total:", total)
}

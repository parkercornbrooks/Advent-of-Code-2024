package day5

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

func (d day) Part1(day int, file string) {
	rules, updates := loadData(day, file)

	total := 0

	for _, update := range updates {
		if rules.orderIsValid(update) {
			total += middlePage(update)
		}
	}
	fmt.Println("Total:", total)
}

func loadData(day int, file string) (Tree, [][]string) {
	isRules := true
	rules := Tree{}
	updates := [][]string{}

	linefn := func(line string) {
		if line == "" {
			isRules = false
		} else if isRules {
			strRule := strings.Split(line, "|")
			rules.addEdge(strRule[0], strRule[1])
		} else {
			strUpdate := strings.Split(line, ",")
			updates = append(updates, strUpdate)
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)
	return rules, updates
}

func middlePage(update []string) int {
	midInd := len(update) / 2
	midVal := update[midInd]
	return utils.MustAtoi(midVal)
}

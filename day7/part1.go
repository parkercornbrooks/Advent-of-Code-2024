package day7

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

const Star = "*"
const Plus = "+"

type Equation struct {
	Result int
	Input  []int
}

// isValid returns the equation result if the equation can be made true
func (e Equation) isValid() int {
	opGroups := generateOpGroups(len(e.Input) - 1)
	for _, g := range opGroups {
		res := e.eval(g)
		if res == e.Result {
			return e.Result
		}
	}
	return 0
}

func (e Equation) eval(ops []string) int {
	res := e.Input[0]
	for i := 0; i < len(e.Input)-1; i++ {
		switch ops[i] {
		case Star:
			res = res * e.Input[i+1]
		case Plus:
			res = res + e.Input[i+1]
		default:
			panic("unknown operator")
		}
	}
	return res
}

func generateOpGroups(count int) [][]string {
	prev := [][]string{
		{Star},
		{Plus},
	}
	next := [][]string{}
	for i := 1; i < count; i++ {
		for _, p := range prev {
			for _, op := range []string{Star, Plus} {
				newS := make([]string, len(p)+1)
				copy(newS, p)
				newS[len(p)] = op
				next = append(next, newS)
			}
		}
		prev = next
		next = [][]string{}
	}
	return prev
}

func (d day) Part1(day int, file string) {
	eqs := readInput(day, file)
	total := 0

	for _, e := range eqs {
		total += e.isValid()
	}

	fmt.Println(total)
}

func readInput(day int, file string) []Equation {
	equations := make([]Equation, 0)

	linefn := func(line string) {
		f := strings.Fields(line)
		res := f[0][:len(f[0])-1]
		input := f[1:]
		inputInts := make([]int, len(input))
		for i, val := range input {
			inputInts[i] = utils.MustAtoi(val)
		}
		e := Equation{
			Result: utils.MustAtoi(res),
			Input:  inputInts,
		}
		equations = append(equations, e)
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return equations
}

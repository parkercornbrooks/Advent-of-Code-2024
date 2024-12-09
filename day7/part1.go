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
const Conc = "|"

type Equation struct {
	Result int
	Input  []int
}

// isValid returns the equation result if the equation can be made true
func (e Equation) isValid(ops []string) int {
	opGroups := generateOpGroups(len(e.Input)-1, ops)
	for _, g := range opGroups {
		res := e.eval(g)
		if res == e.Result {
			return res
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
		case Conc:
			res = utils.MustAtoi(strconv.Itoa(res) + strconv.Itoa(e.Input[i+1]))
		default:
			panic("unknown operator")
		}
	}
	return res
}

var opsCache = make(map[string][][]string)

func makeOpsTag(numOps, count int) string {
	return fmt.Sprintf("%d-%d", numOps, count)
}

func generateOpGroups(count int, ops []string) [][]string {
	tag := makeOpsTag(len(ops), count)
	if cached, ok := opsCache[tag]; ok {
		return cached
	}
	prev := [][]string{}
	for _, op := range ops {
		prev = append(prev, []string{op})
	}
	next := [][]string{}
	for i := 1; i < count; i++ {
		for _, p := range prev {
			for _, op := range ops {
				newS := make([]string, len(p)+1)
				copy(newS, p)
				newS[len(p)] = op
				next = append(next, newS)
			}
		}
		prev = next
		next = [][]string{}
	}
	opsCache[tag] = prev
	return prev
}

func (d day) Part1(day int, file string) int {
	eqs := readInput(day, file)
	total := 0
	ops := []string{Star, Plus}

	for _, e := range eqs {
		total += e.isValid(ops)
	}

	return total
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

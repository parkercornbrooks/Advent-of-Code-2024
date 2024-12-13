package day13

import (
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

const A_SCORE = 3
const B_SCORE = 1

type Machine struct {
	ax, ay, bx, by, px, py int
}

func (m Machine) play() (float64, float64) {
	div := float64(m.by) / float64(m.bx)
	num := (float64(m.py) - (float64(m.px) * div))
	denom := (float64(m.ay) - (float64(m.ax) * div))
	a := num / denom
	b := (float64(m.px) - a*float64(m.ax)) / float64(m.bx)
	return a, b
}

func (d day) Part1(day int, file string) int {
	machines := load(day, file)
	total := 0

	for _, m := range machines {
		a, b := m.play()
		if aInt, bInt, ok := convBoth(a, b); ok {
			total += aInt*A_SCORE + bInt*B_SCORE
		}
	}
	return total
}

func convBoth(a, b float64) (int, int, bool) {
	aInt, aOk := intConv(a)
	bInt, bOk := intConv(b)
	if aOk && bOk {
		return aInt, bInt, true
	}
	return 0, 0, false
}

func intConv(a float64) (int, bool) {
	corr := 1e-8
	aWhole, aFrac := math.Modf(a)
	if aFrac < corr {
		return int(aWhole), true
	} else if aFrac > 1-corr {
		return int(aWhole + 1), true // round up
	}
	return 0, false
}

func load(day int, file string) []Machine {
	machines := []Machine{}

	m := Machine{}

	linefn := func(line string) {
		if strings.HasPrefix(line, "Button A") {
			m.ax, m.ay = parseButton(line)
		} else if strings.HasPrefix(line, "Button B") {
			m.bx, m.by = parseButton(line)
		} else if strings.HasPrefix(line, "P") {
			m.px, m.py = parsePrize(line)
		} else {
			machines = append(machines, m)
			m = Machine{}
		}
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)
	machines = append(machines, m)
	return machines
}

func parseButton(line string) (int, int) {
	fields := strings.Fields(line)
	xS := strings.TrimPrefix(fields[2], "X+")
	xS = strings.TrimSuffix(xS, ",")
	yS := strings.TrimPrefix(fields[3], "Y+")
	return utils.MustAtoi(xS), utils.MustAtoi(yS)
}

func parsePrize(line string) (int, int) {
	fields := strings.Fields(line)
	xS := strings.TrimPrefix(fields[1], "X=")
	xS = strings.TrimSuffix(xS, ",")
	yS := strings.TrimPrefix(fields[2], "Y=")
	return utils.MustAtoi(xS), utils.MustAtoi(yS)
}

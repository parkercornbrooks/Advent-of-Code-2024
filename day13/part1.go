package day13

import (
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

func (m Machine) play() (int, int) {
	a := ((m.by * m.px) - (m.bx * m.py)) / (m.by*m.ax - m.bx*m.ay)
	b := (m.px - a*m.ax) / m.bx
	if (a*m.ay)+(b*m.by) == m.py && (a*m.ax)+(b*m.bx) == m.px {
		return a, b
	}
	return 0, 0
}

func (d day) Part1(day int, file string) int {
	machines := load(day, file)
	total := 0

	for _, m := range machines {
		a, b := m.play()
		total += a*A_SCORE + b*B_SCORE
	}
	return total
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

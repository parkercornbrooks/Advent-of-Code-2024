package day14

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

type Robot struct {
	x, y, vx, vy int
}

func (r *Robot) posAfterSec(s int) (int, int) {
	x := r.x + r.vx*s
	y := r.y + r.vy*s
	return x, y
}

type Area struct {
	w, h          int
	bots          []*Robot
	rowMaxReached bool
}

func (a *Area) reposition(x, y int) (newX, newY int) {
	if x >= 0 {
		newX = x % a.w
	} else {
		newX = (a.w + (x % a.w)) % a.w
	}
	if y >= 0 {
		newY = y % a.h
	} else {
		newY = (a.h + (y % a.h)) % a.h
	}

	return newX, newY
}

func (a *Area) quadrant(x, y int) (int, bool) {
	if y == a.h/2 || x == a.w/2 {
		return 0, false
	}
	q := 0
	if y > a.h/2 {
		q += 2
	}
	if x > a.w/2 {
		q += 1
	}
	return q, true
}

func (d day) Part1(day int, file string) int {
	w, h := 101, 103
	if file == "example.txt" {
		w, h = 11, 7
	}

	a := &Area{w, h, nil, false}

	robots := load(day, file)

	count := map[int]int{}

	for _, r := range robots {
		x, y := r.posAfterSec(100)
		x, y = a.reposition(x, y)
		q, ok := a.quadrant(x, y)
		if ok {
			count[q] += 1
		}
	}

	total := 1

	for _, v := range count {
		total *= v
	}

	return total
}

func load(day int, file string) []*Robot {
	robots := []*Robot{}
	linefn := func(line string) {
		fields := strings.Fields(line)
		pStrip := strings.TrimPrefix(fields[0], "p=")
		point := strings.Split(pStrip, ",")
		vStrip := strings.TrimPrefix(fields[1], "v=")
		vel := strings.Split(vStrip, ",")
		robot := &Robot{
			x:  utils.MustAtoi(point[0]),
			y:  utils.MustAtoi(point[1]),
			vx: utils.MustAtoi(vel[0]),
			vy: utils.MustAtoi(vel[1]),
		}
		robots = append(robots, robot)
	}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return robots
}

package utils

import (
	"bufio"
	"os"
	"strconv"
)

type parseLine func(string)

func ReadInput(filename string, linefn parseLine) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		linefn(sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
}

func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

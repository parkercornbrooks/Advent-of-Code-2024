package utils

import (
	"bufio"
	"os"
)

type parseLine func(string)

type endRead func()

func ReadInput(filename string, linefn parseLine, endfn endRead) {
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
	endfn()
}

func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}

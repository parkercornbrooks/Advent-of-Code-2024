package main

import (
	"flag"
	"fmt"
	"testing"
)

var example = flag.Bool("example", false, "test with the example input")

var testCases = []struct {
	day, part, example, input int
}{
	{
		day:     1,
		part:    1,
		example: 11,
		input:   2742123,
	},
	{
		day:     1,
		part:    2,
		example: 31,
		input:   21328497,
	},
	{
		day:     2,
		part:    1,
		example: 2,
		input:   306,
	},
	{
		day:     2,
		part:    2,
		example: 4,
		input:   366,
	},
	{
		day:     3,
		part:    1,
		example: 161,
		input:   178538786,
	},
	{
		day:     3,
		part:    2,
		example: 48,
		input:   102467299,
	},
	{
		day:     4,
		part:    1,
		example: 18,
		input:   2603,
	},
	{
		day:     4,
		part:    2,
		example: 9,
		input:   1965,
	},
	{
		day:     5,
		part:    1,
		example: 143,
		input:   5087,
	},
	{
		day:     5,
		part:    2,
		example: 123,
		input:   4971,
	},
	{
		day:     6,
		part:    1,
		example: 41,
		input:   4890,
	},
	{
		day:     6,
		part:    2,
		example: 6,
		input:   1995,
	},
	{
		day:     7,
		part:    1,
		example: 3749,
		input:   3312271365652,
	},
	{
		day:     7,
		part:    2,
		example: 11387,
		input:   509463489296712,
	},
	{
		day:     8,
		part:    1,
		example: 14,
		input:   289,
	},
	{
		day:     8,
		part:    2,
		example: 34,
		input:   1030,
	},
	{
		day:     9,
		part:    1,
		example: 1928,
		input:   6337921897505,
	},
	{
		day:     9,
		part:    2,
		example: 2858,
		input:   6362722604045,
	},
	{
		day:     10,
		part:    1,
		example: 36,
		input:   737,
	},
	{
		day:     10,
		part:    2,
		example: 81,
		input:   1619,
	},
	{
		day:     11,
		part:    1,
		example: 55312,
		input:   185205,
	},
	{
		day:     11,
		part:    2,
		example: 65601038650482,
		input:   221280540398419,
	},
	{
		day:     12,
		part:    1,
		example: 1930,
		input:   1477762,
	},
	{
		day:     12,
		part:    2,
		example: 1206,
		input:   923480,
	},
	{
		day:     13,
		part:    1,
		example: 480,
		input:   36870,
	},
	{
		day:     13,
		part:    2,
		example: 875318608908,
		input:   78101482023732,
	},
	{
		day:     14,
		part:    1,
		example: 12,
		input:   214109808,
	},
	{
		day:     14,
		part:    2,
		example: 0,
		input:   7687,
	},
	{
		day:     15,
		part:    1,
		example: 10092,
		input:   1476771,
	},
	{
		day:     15,
		part:    2,
		example: 9021,
		input:   1468005,
	},
	{
		day:     16,
		part:    1,
		example: 7036,
		input:   108504,
	},
	{
		day:     16,
		part:    2,
		example: 0,
		input:   0,
	},
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		name := fmt.Sprintf("Day %d Part %d", tc.day, tc.part)
		file := "input.txt"
		want := tc.input
		if *example {
			file = "example.txt"
			want = tc.example
		}
		t.Run(name, func(t *testing.T) {
			result := run(tc.day, tc.part, file)
			if want != result {
				t.Errorf("Expected %d, got %d instead\n", want, result)
			}
		})
	}
}

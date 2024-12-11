package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	day  int
	part int
	want int
}{
	{
		day:  1,
		part: 1,
		want: 2742123,
	},
	{
		day:  1,
		part: 2,
		want: 21328497,
	},
	{
		day:  2,
		part: 1,
		want: 306,
	},
	{
		day:  2,
		part: 2,
		want: 366,
	},
	{
		day:  3,
		part: 1,
		want: 178538786,
	},
	{
		day:  3,
		part: 2,
		want: 102467299,
	},
	{
		day:  4,
		part: 1,
		want: 2603,
	},
	{
		day:  4,
		part: 2,
		want: 1965,
	},
	{
		day:  5,
		part: 1,
		want: 5087,
	},
	{
		day:  5,
		part: 2,
		want: 4971,
	},
	{
		day:  6,
		part: 1,
		want: 4890,
	},
	{
		day:  6,
		part: 2,
		want: 1995,
	},
	{
		day:  7,
		part: 1,
		want: 3312271365652,
	},
	{
		day:  7,
		part: 2,
		want: 509463489296712,
	},
	{
		day:  8,
		part: 1,
		want: 289,
	},
	{
		day:  8,
		part: 2,
		want: 1030,
	},
	{
		day:  9,
		part: 1,
		want: 6337921897505,
	},
	{
		day:  9,
		part: 2,
		want: 6362722604045,
	},
	{
		day:  10,
		part: 1,
		want: 737,
	},
	{
		day:  10,
		part: 2,
		want: 1619,
	},
	{
		day:  11,
		part: 1,
		want: 185205,
	},
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		name := fmt.Sprintf("Day %d Part %d", tc.day, tc.part)
		t.Run(name, func(t *testing.T) {
			result := run(tc.day, tc.part, "input.txt")
			if tc.want != result {
				t.Errorf("Expected %d, got %d instead\n", tc.want, result)
			}
		})
	}
}

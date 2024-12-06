package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

var testCases = []struct {
	day  int
	part int
	want string
}{
	{
		day:  1,
		part: 1,
		want: "Total: 2742123",
	},
	{
		day:  1,
		part: 2,
		want: "Total: 21328497",
	},
	{
		day:  2,
		part: 1,
		want: "Total safe reports: 306",
	},
	{
		day:  2,
		part: 2,
		want: "Total safe reports: 366",
	},
	{
		day:  3,
		part: 1,
		want: "Total: 178538786",
	},
	{
		day:  3,
		part: 2,
		want: "Total: 102467299",
	},
	{
		day:  4,
		part: 1,
		want: "2603",
	},
	{
		day:  4,
		part: 2,
		want: "1965",
	},
	{
		day:  5,
		part: 1,
		want: "Total: 5087",
	},
	{
		day:  5,
		part: 2,
		want: "Total: 4971",
	},
	{
		day:  6,
		part: 1,
		want: "4890",
	},
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		name := fmt.Sprintf("Day %d Part %d", tc.day, tc.part)
		daycmd := fmt.Sprintf("%d", tc.day)
		partcmd := fmt.Sprintf("%d", tc.part)
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command("go", "run", ".", "-d", daycmd, "-p", partcmd)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			result := strings.Split(string(out), "\n")[1]
			if tc.want != result {
				t.Errorf("Expected %q, got %q instead\n", tc.want, result)
			}
		})
	}
}

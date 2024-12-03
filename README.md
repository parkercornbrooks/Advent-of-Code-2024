# Advent of Code 2024

Attempting in Go this year

## How to Run

run the code using `go run ./dayX`
specify a part using the `-part` flag and a 1 or 2 (default is 1)

```sh
$ go run ./day1 -part=2
```

## Notes

The `utils` package contains a `ReadInput()` function with 3 params:
- filename - this is the input file, which should be saved under the given day and named `input.txt` (e.g. `dayX/input.txt`). These files have been .gitignored as is standard for AoC
- linefn - a function to run based on each line that is read. This is where the input file is read into whatever data format is necessary for the given puzzle. 
- endfn - a function to run after the input file has been fully read, usually the end calculation

I have found this format to be effective for abstracting the reading of the file while providing flexibility for different input structures and problem statements.
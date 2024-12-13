# Advent of Code 2024

Attempting in Go this year

## How to Run

Run the code using `go run .`

Specify the day using the `-d` flag (default is 1)

```sh
$ go run . -d 4
```

Specify a part using the `-p` flag and a 1 or 2 (default is 1)

```sh
$ go run . -p 2
```

Specify the input file using the `-f` flag (default is "input.txt")

```sh
$ go run . -f "example.txt"
```

Example - Run Day 5 Part 2 with the example input
```sh
$ go run . -d 5 -p 2 -f "example.txt"
```

## Notes

### Utils

The `utils` package contains a `ReadInput()` function with 2 params:
- filename - the default is the input file, which should be saved under the given day and named `input.txt` (e.g. `dayX/input.txt`). These files have been .gitignored as is standard for AoC. Alternatively it will read the passed `-f` flag under the same directory (e.g. `dayX/myOtherFile.txt`)
- linefn - a function to run based on each line that is read. This is where the input file is read into whatever data format is necessary for the given puzzle. 

I have found this format to be effective for abstracting the reading of the file while providing flexibility for different input structures and problem statements.

### Testing

Added a test suite so that if I change the utils I can confirm that the earlier runs still function as expected.

To use it, run either:

```sh
$ go test
$ go test -v
```

Since my specific input files are not available publicly, the `-example` flag can be applied to run the test suite against the example text files in the same manner:

```sh
$ go test -example
$ go test -v -example
```

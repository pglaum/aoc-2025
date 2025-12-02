```go
package main

import (
	"flag"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	sample := flag.Bool("sample", false, "use sample input")
	flag.Parse()

	if *sample {
		filename = "sample.txt"
	}

	data, _ := os.ReadFile(filename)

	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\r\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
}

func part2(lines []string) {
}
```

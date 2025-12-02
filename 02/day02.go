package main

import (
	"flag"
	"os"
	"strconv"
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
	ranges := strings.Split(input, ",")

	part1(ranges)
	part2(ranges)
}

func part1(ranges []string) {
	count := 0
	for r := range ranges {
		ints := strings.Split(ranges[r], "-")
		from, _ := strconv.Atoi(ints[0])
		to, _ := strconv.Atoi(ints[1])

		for i := from; i <= to; i++ {
			s := strconv.Itoa(i)
			l := len(s)
			if l%2 != 0 {
				continue
			}

			if s[:l/2] == s[l/2:] {
				count += i
			}
		}
	}
	println("Part 1:", count)
}

func part2(ranges []string) {
	count := 0
	for r := range ranges {
		ints := strings.Split(ranges[r], "-")
		from, _ := strconv.Atoi(ints[0])
		to, _ := strconv.Atoi(ints[1])

		for i := from; i <= to; i++ {
			s := strconv.Itoa(i)
			l := len(s)
			maxWidth := l / 2

			for width := 1; width <= maxWidth; width++ {
				if l%width != 0 {
					continue
				}

				first := s[:width]
				matched := true

				for j := width; j < l; j += width {
					if s[j:j+width] == first {
						continue
					}
					matched = false
					break
				}

				if matched {
					count += i
					break
				}
			}
		}
	}
	println("Part 2:", count)
}

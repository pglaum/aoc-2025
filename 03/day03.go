package main

import (
	"flag"
	"math"
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
	count := 0
	for _, line := range lines {
		tens := 0
		idx := 0
		for i, char := range line[:len(line)-1] {
			num := int(char - '0')
			if num > tens {
				tens = num
				idx = i
			}
		}

		ones := 0
		for _, char := range line[idx+1:] {
			num := int(char - '0')
			if num > ones {
				ones = num
			}
		}

		count += tens*10 + ones
	}
	println("Part 1:", count)
}

func part2(lines []string) {
	count := 0
	for _, line := range lines {
		res := 0
		next := 0
		for di := range 12 {
			d := 11 - di
			high := 0
			start := next
			for i, char := range line[start : len(line)-d] {
				num := int(char - '0')
				if num > high {
					high = num
					next = start + i + 1
				}
			}
			res += int(math.Pow10(int(d))) * high
		}

		count += res
	}
	println("Part 2:", count)
}

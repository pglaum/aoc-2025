package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\r\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	dial := 50
	password := 0
	for _, line := range lines {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		if dir == 'L' {
			dial -= steps
		} else {
			dial += steps
		}

		dial %= 100
		if dial == 0 {
			password += 1
		}
	}
	println("Part 1:", password)
}

func part2(lines []string) {
	dial := 50
	password := 0
	for _, line := range lines {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		sign := 1
		if dir == 'L' {
			sign = -1
		}

		for range steps {
			dial += sign
			dial %= 100
			if dial == 0 {
				password += 1
			}
		}
	}
	println("Part 2:", password)
}

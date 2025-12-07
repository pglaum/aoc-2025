package main

import (
	"flag"
	"os"
	"slices"
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
	tachs := []int{}
	for i, c := range lines[0] {
		if c == 'S' {
			tachs = []int{i}
			break
		}
	}

	splits := 0
	for _, line := range lines[1:] {
		newTachs := []int{}
		for _, t := range tachs {
			switch line[t] {
			case '^':
				splits += 1
				if !slices.Contains(newTachs, t-1) {
					newTachs = append(newTachs, t-1)
				}
				if !slices.Contains(newTachs, t+1) {
					newTachs = append(newTachs, t+1)
				}
			case '.':
				if !slices.Contains(newTachs, t) {
					newTachs = append(newTachs, t)
				}
			default:
				println("Unknown char", line[t])
			}
		}
		tachs = newTachs
	}
	println("Part 1:", splits)
}

func part2(lines []string) {
	tachs := []int{}

	for i, c := range lines[0] {
		if c == 'S' {
			tachs = []int{i}
			break
		}
	}

	ol := map[int]int{
		tachs[0]: 1,
	}

	splits := 0
	for _, line := range lines[1:] {
		newTachs := []int{}
		newOl := map[int]int{}
		for _, t := range tachs {
			switch line[t] {
			case '^':
				splits += 1
				if !slices.Contains(newTachs, t-1) {
					newTachs = append(newTachs, t-1)
				}
				if !slices.Contains(newTachs, t+1) {
					newTachs = append(newTachs, t+1)
				}
				newOl[t-1] += ol[t]
				newOl[t+1] += ol[t]
			case '.':
				if !slices.Contains(newTachs, t) {
					newTachs = append(newTachs, t)
				}
				newOl[t] += ol[t]
			default:
				println("Unknown char", line[t])
			}
		}

		ol = newOl
		tachs = newTachs
	}

	count := 0
	for _, v := range ol {
		count += v
	}

	println("Part 2:", count)
}

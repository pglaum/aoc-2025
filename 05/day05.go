package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
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
	parts := strings.Split(input, "\r\n\r\n")
	db := strings.Split(parts[0], "\r\n")
	lines := strings.Split(parts[1], "\r\n")

	ranges := [][2]int{}
	for _, line := range db {
		var start, end int
		fmt.Sscanf(line, "%d-%d", &start, &end)
		ranges = append(ranges, [2]int{start, end})
	}

	ingredients := []int{}
	for _, line := range lines {
		var ingredient int
		fmt.Sscanf(line, "%d", &ingredient)
		ingredients = append(ingredients, ingredient)
	}

	part1(ranges, ingredients)
	part2(ranges)
}

func part1(ranges [][2]int, ingredients []int) {
	count := 0
	for _, ingredient := range ingredients {
		fresh := false
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				fresh = true
				break
			}
		}
		if fresh {
			count += 1
		}
	}
	println("Part 1:", count)
}

func part2(ranges [][2]int) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	loops := 0
	for {
		changes := 0
		newRanges := [][2]int{}
		for i, r := range ranges {
			found := false
			for j, r2 := range newRanges {
				if r[0] >= r2[0] && r[0] <= r2[1] && r[1] >= r2[1] {
					newRanges[j][1] = r[1]
					found = true
					break
				}
				if r[1] >= r2[0] && r[1] <= r2[1] && r[0] <= r2[0] {
					newRanges[j][0] = r[0]
					found = true
					break
				}
				if r[0] >= r2[0] && r[1] <= r2[1] {
					found = true
					break
				}
			}
			if !found {
				newRanges = append(newRanges, ranges[i])
			} else {
				changes += 1
			}
		}

		copy(ranges, newRanges)
		ranges = ranges[:len(newRanges)]

		if changes == 0 {
			break
		}
		loops += 1
	}

	count := 0
	for _, r := range ranges {
		count += r[1] - r[0] + 1
	}
	println("Part 2:", count)
}

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

func getAdjacent(rolls [][]int, x, y int) int {
	count := 0
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if newY >= 0 && newY < len(rolls) && newX >= 0 && newX < len(rolls[newY]) {
			if rolls[newY][newX] > 0 {
				count++
			}
		}
	}
	return count
}

func part1(lines []string) {
	rolls := [][]int{}
	for i, line := range lines {
		rolls = append(rolls, []int{})
		for _, char := range line {
			paperRoll := 0
			if char == '@' {
				paperRoll = 1
			}
			rolls[i] = append(rolls[i], paperRoll)
		}
	}

	count := 0
	for y := range rolls {
		for x := range rolls[y] {
			if rolls[y][x] == 0 {
				continue
			}

			adj := getAdjacent(rolls, x, y)
			if adj < 4 {
				rolls[y][x] = 2
				count++
			}
		}
	}
	println("Part 1:", count)
}

func part2(lines []string) {
	rolls := [][]int{}
	for i, line := range lines {
		rolls = append(rolls, []int{})
		for _, char := range line {
			paperRoll := 0
			if char == '@' {
				paperRoll = 1
			}
			rolls[i] = append(rolls[i], paperRoll)
		}
	}

	rounds := 0
	totalCount := 0

	for {
		count := 0

		for y := range rolls {
			for x := range rolls[y] {
				if rolls[y][x] == 0 {
					continue
				}

				adj := getAdjacent(rolls, x, y)
				if adj < 4 {
					rolls[y][x] = 2
					count++
				}
			}
		}
		for y := range rolls {
			for x := range rolls[y] {
				if rolls[y][x] == 2 {
					rolls[y][x] = 0
				}
			}
		}

		totalCount += count
		if count == 0 {
			break
		}
		rounds++
	}
	println("Part 2:", totalCount)
}

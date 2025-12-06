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

	lines := strings.Split(string(data), "\r\n")
	lines = lines[:len(lines)-1]

	values := [][]int{}
	for i, line := range lines[:len(lines)-1] {
		values = append(values, []int{})
		fields := strings.FieldsSeq(line)

		for f := range fields {
			num, _ := strconv.Atoi(f)
			values[i] = append(values[i], num)
		}
	}

	operators := []string{}
	fields := strings.FieldsSeq(lines[len(lines)-1])
	for f := range fields {
		operators = append(operators, f)
	}

	part1(values, operators)
	part2(lines)
}

func part1(values [][]int, operators []string) {
	count := 0
	for i := range len(operators) {
		switch operators[i] {
		case "+":
			c := 0
			for j := range len(values) {
				c += values[j][i]
			}
			count += c
		case "*":
			c := 1
			for j := range len(values) {
				c = c * values[j][i]
			}
			count += c
		}
	}
	println("Part 1:", count)
}

func part2(lines []string) {
	count := 0
	operator := ""
	values := [6]string{}
	newColStart := 0

	for j := range len(lines[0]) {

		newCol := true
		for i := range len(lines) {
			char := lines[i][j]
			if char == ' ' {
				continue
			}
			newCol = false
			if i == len(lines)-1 {
				operator = string(char)
				continue
			}
			values[j-newColStart] += string(char)
		}
		if j == len(lines[0])-1 {
			newCol = true
		}

		if newCol {
			var c int
			switch operator {
			case "+":
				c = 0
				for _, v := range values {
					val, _ := strconv.Atoi(v)
					c += val
				}
				count += c
			case "*":
				c = 1
				for _, v := range values {
					val, _ := strconv.Atoi(v)
					if val == 0 {
						continue
					}
					c = c * val
				}
				count += c
			}

			operator = ""
			values = [6]string{}
			newColStart = j + 1
		}
	}
	println("Part 2:", count)
}

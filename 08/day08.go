package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strings"
)

var (
	CONNECT_COUNT = 1000
)

func main() {
	filename := "input.txt"
	sample := flag.Bool("sample", false, "use sample input")
	flag.Parse()

	if *sample {
		filename = "sample.txt"
		CONNECT_COUNT = 10
	}

	data, _ := os.ReadFile(filename)

	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\r\n")

	vectors := []Vec3{}
	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		vectors = append(vectors, Vec3{x, y, z})
	}

	part1(vectors)
	part2(vectors)
}

type Vec3 struct {
	x int
	y int
	z int
}

func (v *Vec3) Dist(other Vec3) float64 {
	return math.Sqrt(
		math.Pow(float64(v.x)-float64(other.x), 2) +
			math.Pow(float64(v.y)-float64(other.y), 2) +
			math.Pow(float64(v.z)-float64(other.z), 2),
	)
}

func (v *Vec3) LessThan(other Vec3) bool {
	if v.x != other.x {
		return v.x < other.x
	}
	if v.y != other.y {
		return v.y < other.y
	}
	return v.z < other.z
}

func part1(vectors []Vec3) {
	distances := map[[2]int]float64{}
	for i := range vectors {
		for j := range vectors {
			if i == j {
				continue
			}
			dist := vectors[i].Dist(vectors[j])
			if i < j {
				distances[[2]int{i, j}] = dist
			} else {
				distances[[2]int{j, i}] = dist
			}
		}
	}

	keys := [][2]int{}
	for k := range distances {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return distances[keys[i]] < distances[keys[j]]
	})

	connections := [][2]int{}
	for i := range CONNECT_COUNT {
		connections = append(connections, keys[i])
	}

	circuits := [][]int{}
	for _, conn := range connections {
		found := false
		firstI := -1
		for i, c := range circuits {
			if slices.Contains(c, conn[0]) || slices.Contains(c, conn[1]) {
				if found {
					for _, val := range circuits[i] {
						if !slices.Contains(circuits[firstI], val) {
							circuits[firstI] = append(circuits[firstI], val)
						}
					}
					circuits[i] = []int{}
					continue
				} else {
					found = true
					firstI = i
				}
				if !slices.Contains(circuits[i], conn[0]) {
					circuits[i] = append(circuits[i], conn[0])
				}
				if !slices.Contains(circuits[i], conn[1]) {
					circuits[i] = append(circuits[i], conn[1])
				}
			}
		}
		if !found {
			circuits = append(circuits, []int{conn[0], conn[1]})
		}
	}

	lens := []int{}
	for _, c := range circuits {
		if len(c) > 0 {
			lens = append(lens, len(c))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lens)))

	count := 1
	for i := range 3 {
		count *= lens[i]
	}
	println("Part 1:", count)
}

func part2(vectors []Vec3) {
	distances := map[[2]int]float64{}
	for i := range vectors {
		for j := range vectors {
			if i == j {
				continue
			}
			dist := vectors[i].Dist(vectors[j])
			if i < j {
				distances[[2]int{i, j}] = dist
			} else {
				distances[[2]int{j, i}] = dist
			}
		}
	}

	keys := [][2]int{}
	for k := range distances {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return distances[keys[i]] < distances[keys[j]]
	})

	circuits := [][]int{}
	for _, conn := range keys {
		found := false
		firstI := -1
		for i, c := range circuits {
			if slices.Contains(c, conn[0]) || slices.Contains(c, conn[1]) {
				if found {
					for _, val := range circuits[i] {
						if !slices.Contains(circuits[firstI], val) {
							circuits[firstI] = append(circuits[firstI], val)
						}
					}
					circuits[i] = []int{}
					continue
				} else {
					found = true
					firstI = i
				}
				if !slices.Contains(circuits[i], conn[0]) {
					circuits[i] = append(circuits[i], conn[0])
				}
				if !slices.Contains(circuits[i], conn[1]) {
					circuits[i] = append(circuits[i], conn[1])
				}
			}
		}
		if !found {
			circuits = append(circuits, []int{conn[0], conn[1]})
		}

		cleanCircuits := [][]int{}
		for _, c := range circuits {
			if len(c) > 0 {
				cleanCircuits = append(cleanCircuits, c)
			}
		}
		if len(cleanCircuits[0]) >= len(vectors) {
			fmt.Printf("Part 2: %d", vectors[conn[0]].x*vectors[conn[1]].x)
			break
		}
	}
}

package main

import (
	aoc "aoc/utils"
	"fmt"
	"sort"
	"strings"
)

var (
	si        [][]int
	colMax    int
	rowMax    int
	lowPoints map[[2]int]struct{}
)

func p1() {
	var risk int
	for row := range si {
		for col := range si[row] {
			height := si[row][col]
			if height == 9 {
				continue
			}
			skip := false
			for _, adj := range adjacent(row, col) {
				if height > si[adj[0]][adj[1]] {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			lowPoints[[2]int{row, col}] = struct{}{}
			risk += 1 + height
		}
	}
	fmt.Println(risk)
}

func adjacent(row, col int) (adj [][2]int) {
	if row > 0 {
		adj = append(adj, [2]int{row - 1, col})
	}
	if row < rowMax {
		adj = append(adj, [2]int{row + 1, col})
	}
	if col > 0 {
		adj = append(adj, [2]int{row, col - 1})
	}
	if col < colMax {
		adj = append(adj, [2]int{row, col + 1})
	}
	return
}

func flow(row, col int) (lp [2]int) {
	adjCells := adjacent(row, col)
	for _, adj := range adjCells {
		// Check for low points
		if _, islp := lowPoints[adj]; islp {
			return adj
		}
	}
	for _, adj := range adjCells {
		if si[adj[0]][adj[1]] < si[row][col] {
			return flow(adj[0], adj[1])
		}
	}
	panic("Error in flow!")
}

// Could store the number of times we flow
// before finding the low point (and the positions)
// and add that, that way we avoid looping over
// every position
func p2() {
	basins := make(map[[2]int]int) // lowpoint: size
	for row := range si {
		for col := range si[row] {
			if si[row][col] == 9 {
				continue
			}
			if _, islp := lowPoints[[2]int{row, col}]; islp {
				basins[[2]int{row, col}]++
				continue
			}
			lp := flow(row, col)
			basins[lp]++
		}
	}
	sizes := []int{}
	for _, size := range basins {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	fmt.Println(sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1])
}

func main() {
	ss := aoc.GetLines("input.txt")
	for _, s := range ss {
		si = append(si, aoc.SStoSI(strings.Split(s, "")))
	}
	rowMax = len(si) - 1
	colMax = len(si[0]) - 1
	lowPoints = make(map[[2]int]struct{})
	p1() // 607
	p2() // 900864
}

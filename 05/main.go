package main

import (
	aoc "aoc/utils"
	"fmt"
)

var (
	lines   [][4]int
	covered map[Pos]int
)

type Pos struct {
	x, y int
}

func load(fname string) {
	for _, s := range aoc.GetLines(fname) {
		var line [4]int
		fmt.Sscanf(s, "%d,%d -> %d,%d", &line[0], &line[1], &line[2], &line[3])
		lines = append(lines, line)
	}
}

func solve(p2 bool) {
	covered = make(map[Pos]int)
	for _, l := range lines {
		x1, y1 := l[0], l[1]
		x2, y2 := l[2], l[3]
		// Check horizontal and vertical lines
		if x1 == x2 {
			min, max := aoc.MinMax(y1, y2)
			for i := min; i <= max; i++ {
				covered[Pos{x1, i}]++
			}
		}
		if y1 == y2 {
			min, max := aoc.MinMax(x1, x2)
			for i := min; i <= max; i++ {
				covered[Pos{i, y1}]++
			}
		}
		if p2 {
			// Check for diagonals
			if aoc.Abs(x2-x1) != aoc.Abs(y2-y1) {
				continue
			}
			xMin, xMax := aoc.MinMax(x1, x2)
			yMin, yMax := aoc.MinMax(y1, y2)

			for x := xMin; x <= xMax; x++ {
				for y := yMin; y <= yMax; y++ {
					dx, dy := x1-x, y1-y
					if aoc.Abs(dx) == aoc.Abs(dy) {
						covered[Pos{x, y}]++
					}
				}
			}
		}
	}

	res := 0
	for _, count := range covered {
		if count > 1 {
			res++
		}
	}
	fmt.Println(res)
}

func main() {
	load("input.txt")
	solve(false) // 5585
	solve(true)  // 17193
}

package main

import (
	aoc "aoc/utils"
	"fmt"
)

var (
	si         [][]int
	dirVecs    [8][2]int
	flashed    map[[2]int]struct{}
	flashCount int
)

func inBounds(pos [2]int) bool {
	row, col := pos[0], pos[1]
	minRow, maxRow := 0, len(si)-1
	minCol, maxCol := 0, len(si[0])-1
	if row > maxRow || row < minRow || col > maxCol || col < minCol {
		return false
	}
	return true
}

func adjacent(pos [2]int) (adj [][2]int) {
	for _, dir := range dirVecs {
		adjPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if inBounds(adjPos) {
			adj = append(adj, adjPos)
		}
	}
	return
}

func flash(pos [2]int) (change bool) {
	if _, in := flashed[pos]; in {
		return false
	}
	flashed[pos] = struct{}{}
	flashCount++
	for _, adj := range adjacent(pos) {
		row, col := adj[0], adj[1]
		si[row][col]++
		if si[row][col] > 9 {
			flash(adj)
		}
	}
	return true
}

func step() {
	flashed = make(map[[2]int]struct{})
	for row := range si {
		for col := range si[row] {
			si[row][col]++
		}
	}
	for row := range si {
		for col := range si[row] {
			if si[row][col] > 9 {
				flash([2]int{row, col})
			}
		}
	}
	for pos := range flashed {
		si[pos[0]][pos[1]] = 0
	}
}

func p1() {
	for i := 0; i < 100; i++ {
		step()
	}
	fmt.Println(flashCount)
	flashCount = 0
}

func p2() {
	i := 0
	for {
		step()
		i++
		if len(flashed) == len(si)*len(si[0]) {
			break
		}
	}
	fmt.Println(i)
}

func main() {
	si = aoc.GetDigitGrid("input.txt")
	// Deep copy
	bck := make([][]int, len(si))
	for row := range si {
		bck[row] = make([]int, len(si[row]))
		copy(bck[row], si[row])
	}
	p1() // 1773
	p2() // 494
	si = bck
}

func init() {
	dirVecs = [8][2]int{
		{0, 1}, {1, 0},
		{0, -1}, {-1, 0},
		{1, -1}, {-1, 1},
		{1, 1}, {-1, -1},
	}
}

func test(a string) {

}

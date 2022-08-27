package main

import (
	aoc "aoc/utils"
	"fmt"
	"strings"
	"time"
)

type Coord [2]int

const (
	ROW = iota
	COL
)

const (
	AXIS = iota
	VALUE
)

var (
	dots  map[Coord]bool
	folds []Coord
)

func printGrid() {
	maxRow, maxCol := 0, 0
	minRow, minCol := -1, -1
	for dot := range dots {
		if dot[ROW] > maxRow {
			maxRow = dot[ROW]
		}
		if dot[COL] > maxCol {
			maxCol = dot[COL]
		}
		if -1 == minRow || dot[ROW] < minRow {
			minRow = dot[ROW]
		}
		if -1 == minCol || dot[COL] < minCol {
			minCol = dot[COL]
		}
	}
	for col := minCol; col <= maxCol; col++ {
		for row := minRow; row <= maxRow; row++ {
			char := " "
			pos := Coord{row, col}
			if dots[pos] {
				char = "@"
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func makeFold(axis, value int) {
	for pos := range dots {
		if pos[axis] > value {
			newPos := pos
			newPos[axis] = value - (pos[axis] - value)
			delete(dots, pos)
			dots[newPos] = true
		}
	}
}

func solve(p1 bool, label string) {
	defer aoc.TimeTrack(time.Now(), label)
	if p1 {
		fold := folds[0]
		makeFold(fold[AXIS], fold[VALUE])
		fmt.Println(len(dots))
		return
	}
	for _, fold := range folds {
		makeFold(fold[AXIS], fold[VALUE])
	}
	printGrid()
}

func load() {
	defer aoc.TimeTrack(time.Now(), "Loading")
	secs := aoc.GetSections("input.txt", "")
	dots = make(map[Coord]bool)
	for _, line := range secs[0] {
		dot := strings.Split(line, ",")
		pos := Coord{aoc.ToInt(dot[ROW]), aoc.ToInt(dot[COL])}
		dots[pos] = true
	}
	for _, line := range secs[1] {
		instr := strings.Split(line, "=")
		axis := ROW
		if 'y' == instr[0][len(instr[0])-1] {
			axis = COL
		}
		folds = append(folds, Coord{axis, aoc.ToInt(instr[1])})
	}
}

func main() {
	load()
	solve(true, "Part one")  // 729
	solve(false, "Part two") // RGZLBHFP
}

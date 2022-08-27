package main

import (
	aoc "aoc/utils"
	"fmt"
	"strings"
)

const (
	BOARD_LEN = 5 // Square
)

var (
	boards  []Board
	toDraw  []int
	drawPos int
)

type Board struct {
	vals   []int // Let's try a flat slice
	marked []bool
}

func (b *Board) isVictorious() (win bool) {
	// Check the rows
	for row := 0; row < BOARD_LEN; row++ {
		hasRow, hasCol := true, true
		for cell := 0; cell < BOARD_LEN; cell++ {

			ridx := cell + (row * BOARD_LEN)
			if !b.marked[ridx] {
				hasRow = false
			}

			cidx := row + (cell * BOARD_LEN)
			if !b.marked[cidx] {
				hasCol = false
			}

		}
		if hasRow || hasCol {
			return true
		}
	}
	return false
}

// Sum of all unmarked numbers * last drawn number
func (b *Board) calculateScore() (score int) {
	unmarked := 0
	for x := range b.vals {
		if !b.marked[x] {
			unmarked += b.vals[x]
		}
	}
	return unmarked * toDraw[drawPos]
}

func load(fname string) {
	sections := aoc.GetSections(fname, "")
	for i, sec := range sections {
		// Load the numbers to draw
		if i == 0 {
			for _, row := range sec {
				for _, num := range strings.Split(row, ",") {
					toDraw = append(toDraw, aoc.ToInt(num))
				}
			}
			continue
		}
		// Load the board
		b := Board{}
		b.marked = make([]bool, BOARD_LEN*BOARD_LEN)
		for _, row := range sec {
			for _, num := range strings.Fields(row) {
				b.vals = append(b.vals, aoc.ToInt(num))
			}
		}
		boards = append(boards, b)
	}
}

func drawNumber() (winner int, win bool) {
	drawn := toDraw[drawPos]
	for bidx, b := range boards {
		for i := range b.vals {
			if b.vals[i] == drawn {
				b.marked[i] = true
				if b.isVictorious() {
					return bidx, true
				}
			}
		}
	}
	drawPos++
	return -1, false
}

func p1() {
	var winner int
	for win := false; !win; {
		if winner, win = drawNumber(); win {
			fmt.Println(boards[winner].calculateScore())
		}
	}
}

func p2() {
	var winners []Board
	for len(boards) > 0 {
		var winner int
		for win := false; !win; {
			if winner, win = drawNumber(); win {
				// Append to the winners
				winners = append(winners, boards[winner])
				// Remove the board
				boards[winner] = boards[len(boards)-1]
				boards[len(boards)-1] = Board{}
				boards = boards[:len(boards)-1]
			}
		}
	}
	fmt.Println(winners[len(winners)-1].calculateScore())
}

func main() {
	load("input.txt")
	p1() // 60368
	p2() // 17435
}

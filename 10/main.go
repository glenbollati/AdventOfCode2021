package main

import (
	aoc "aoc/utils"
	"fmt"
	"sort"
)

type TokenType int

var (
	ss      []string
	points  map[rune]int
	points2 map[rune]int
	pairs   map[rune]rune
)

func checkLine(s string) (score int, stack []rune) {
	for _, c := range s {
		// Check if its a closing symbol and compare with
		// current expected closing symbol
		if isClosing(c) {
			lastOpen := stack[len(stack)-1]
			expected := pairs[lastOpen]

			if c == expected {
				stack = stack[:len(stack)-1]
				continue
			}
			return points[c], []rune{}
		}
		// We have an opening symbol, add it to the stack
		stack = append(stack, c)
	}
	return 0, stack
}

func p1() {
	var total int
	for _, s := range ss {
		// If stack is not empty, line is incomplete
		score, _ := checkLine(s)
		total += score
	}
	fmt.Println(total)
}

func p2() {
	var scores []int
	for _, s := range ss {
		lineScore := 0
		// If stack of leftover open symbols is not empty, line is incomplete
		_, unmatched := checkLine(s)
		if len(unmatched) > 0 {
			// Walk through the leftover open symbols from back to front, calculating
			// the score for the associated closing symbol
			for i := len(unmatched) - 1; i >= 0; i-- {
				lineScore *= 5
				lineScore += points2[unmatched[i]]
			}
			scores = append(scores, lineScore)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func isClosing(r rune) bool {
	return r == ')' || r == ']' || r == '}' || r == '>'
}

func main() {
	points = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	pairs = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	points2 = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	ss = aoc.GetLines("input.txt")
	p1() // 392367
	p2() // 2192104158
}

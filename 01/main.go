package main

import (
	aoc "aoc/utils"
	"fmt"
)

var ss []string
var si []int

/*
 * Note for part 2
 * For every 4 elements we compare, the middle two in the set are common
 * so it is sufficient to compare the first and last
 */

func solve(p2 bool) (acc int) {
	if !p2 {
		for i := 1; i < len(si); i++ {
			if si[i] > si[i-1] {
				acc++
			}
		}
		fmt.Println(acc)
		return
	}
	for i := 3; i < len(si); i++ {
		if si[i-3] < si[i] {
			acc++
		}
	}
	fmt.Println(acc)
	return
}

func main() {
	ss = aoc.GetLines("input.txt")
	si = aoc.SStoSI(ss)
	solve(false) // 1195
	solve(true)  // 1235
}

package main

import (
	aoc "aoc/utils"
	"fmt"
	"strings"
)

var ss []string

/*
 * Note for part 2
 * As far as part 1 calculations are concerned the additional "aim" variable
 * is equivalent to "depth", so we can just use "aim"
 * and assign depth to it at the end
 */

func solve(p2 bool) {
	var h, d, a int
	for _, s := range ss {
		x := strings.Split(s, " ")
		val := aoc.ToInt(x[1])
		switch x[0] {
		case "forward":
			h += val
			if p2 {
				d += (a * val)
			}
		case "up":
			a -= val
		case "down":
			a += val
		}
	}
	if !p2 {
		d = a
	}
	fmt.Println(h * d)
}

func main() {
	ss = aoc.GetLines("input.txt")
	solve(false) // 1813801
	solve(true)  // 1960569556
}

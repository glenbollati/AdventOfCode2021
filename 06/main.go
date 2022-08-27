package main

import (
	aoc "aoc/utils"
	"fmt"
	"os"
	"strings"
)

// [9]int holds timer counts 0 -> 8
func load(fname string) (timers [9]int) {
	f, _ := os.ReadFile(fname)
	f = f[:len(f)-1] // strip final newline
	for _, x := range strings.Split(string(f), ",") {
		timers[aoc.ToInt(x)]++
	}
	return
}

func solve(timers [9]int, lastDay int) {
	for day := 0; day < lastDay; day++ {
		newTimers := [9]int{}
		for i := range timers[:len(timers)-1] {
			// Downshift all values into new array
			newTimers[i] = timers[i+1]
		}
		// each 0 becomes a 6 and yields a new 8
		newTimers[6] += timers[0]
		newTimers[8] = timers[0]
		timers = newTimers
	}
	var acc int
	for _, count := range timers {
		acc += count
	}
	fmt.Println(acc)
}

func main() {
	timers := load("input.txt")
	solve(timers, 80)
	solve(timers, 256)
}

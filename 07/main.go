package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"sort"
	"time"
)

var (
	si    []int
	fname string
)

func median(si []int) int {
	// if odd, middle value
	if len(si)%2 != 0 {
		return si[(len(si)+1)/2]
	}
	// if even, average of middle two values
	return (si[(len(si)/2)-1] + si[(len(si)/2)]) / 2
}

func mean(si []int) float64 {
	var sum int
	for _, x := range si {
		sum += x
	}
	return float64(sum) / float64(len(si))
}

func p1() {
	defer aoc.TimeTrack(time.Now(), "Part one")
	var fuel int
	m := median(si)
	for _, x := range si {
		fuel += aoc.Abs(m - x)
	}
	fmt.Println(fuel)
}

// The average may not be an integer, in which case
// we need the nearest integers (floor and ceiling)
// we then find the fuel cost for each and take the
// lowest
func p2() {
	defer aoc.TimeTrack(time.Now(), "Part two")
	var fuelFloor, fuelCeil int

	avg := mean(si)
	floor, ceil := int(math.Floor(avg)), int(math.Ceil(avg))

	for _, x := range si {
		f, c := aoc.Abs(floor-x), aoc.Abs(ceil-x)
		fuelFloor += (f * (1 + f) / 2) // Gauss (arithmetic) summation
		fuelCeil += (c * (1 + c) / 2)
	}

	fmt.Println(aoc.Min(fuelFloor, fuelCeil))
}

func load() {
	defer aoc.TimeTrack(time.Now(), "Loading")
	for _, s := range aoc.SplitFile(fname, ",") {
		si = append(si, aoc.ToInt(s))
	}
	sort.Ints(si) // for median calculation in part1
}

func main() {
	fname = "test.txt"
	load()
	p1() // 329389
	p2() // 86397080
}

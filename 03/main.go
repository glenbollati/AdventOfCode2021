package main

import (
	aoc "aoc/utils"
	"fmt"
)

var ss []string

func mclc(ss []string) (mc, lc string) {
	for i := range ss[0] {
		var ones int
		for _, s := range ss {
			if s[i] == '1' {
				ones++
			}
		}
		// n. zeroes = tot - n. ones
		if ones >= len(ss)-ones {
			mc += "1"
			lc += "0"
			continue
		}
		mc += "0"
		lc += "1"
	}
	return
}

func rating(oxy bool) string {
	bins := ss
	for len(bins) > 1 {
		for i := range bins[0] {
			tmp := []string{}
			// outer bounds check doesnt protect us here
			if len(bins) == 1 {
				break
			}
			mc, lc := mclc(bins)
			for _, b := range bins {
				// oxygen => MOST  common bit must match
				// co2    => LEAST common bit must match
				if (oxy && b[i] == mc[i]) || (!oxy && b[i] == lc[i]) {
					tmp = append(tmp, b)
				}
			}
			bins = tmp
		}
	}
	return bins[0]
}

func ans(a, b string) {
	fmt.Println(aoc.BinStrToI64(a) * aoc.BinStrToI64(b))
}

func p1() {
	mc, lc := mclc(ss)
	ans(mc, lc)
}

func p2() {
	oxy := rating(true)
	co2 := rating(false)
	ans(oxy, co2)
}

func main() {
	inputf := "input.txt"
	ss = aoc.GetLines(inputf)
	p1() // 2250414
	p2() // 6085575
}

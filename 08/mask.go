package main

import (
	aoc "aoc/utils"
	"fmt"
	"strings"
	"time"
)

var (
	ss       []string
	sigMap   map[int]int
	runeMap  map[rune]int
	knownLen map[int]int
)

func p1() {
	defer aoc.TimeTrack(time.Now(), "Part one")
	var acc int
	for _, s := range ss {
		output := strings.Split(s, " | ")[1]
		for _, digit := range strings.Split(output, " ") {
			if _, known := knownLen[len(digit)]; known {
				acc++
			}
		}
	}
	fmt.Println(acc)
}

func encode(s string) (set int) {
	for _, r := range s {
		set |= (1 << runeMap[r])
	}
	return
}

func contains(full, sub int) bool {
	return full&sub == sub
}

func bitDiff(a, b int) (count int) {
	xor := a ^ b
	for xor > 0 {
		xor = xor & (xor - 1)
		count++
	}
	return
}

func decode(signals string) (digits map[int]int) {
	digits = make(map[int]int) // digit: bitset
	sigs := strings.Split(signals, " ")

	// First collect the digits of unique length
	toCheck := []int{}
	for i, sig := range sigs {
		if digit, known := knownLen[len(sig)]; known {
			ds := encode(sig)
			digits[digit] = ds
			continue
			// THIS WORKS
		}
		toCheck = append(toCheck, i)
	}

	// TODO
	// Missing: 5, 0, 6

	// Now find the rest
	for i := range toCheck {
		sig := sigs[i]
		ds := encode(sig)
		if len(sig) == 6 {
			if contains(ds, digits[4]) {
				digits[9] = ds
				continue
			}
			if contains(ds, digits[1]) {
				digits[0] = ds
				continue
			}
			digits[6] = ds
		}
		// len == 5
		if contains(ds, digits[1]) {
			digits[3] = ds
			continue
		}
		if 2 == bitDiff(ds, digits[4]) {
			digits[5] = ds
			continue
		}
		digits[2] = ds
	}
	return
}

func p2() {
	defer aoc.TimeTrack(time.Now(), "Part two")
	var acc int
	for _, s := range ss {
		tmp := strings.Split(s, " | ")
		signals, output := tmp[0], tmp[1]

		keys := decode(signals)
		for dig, bs := range keys {
			fmt.Printf("Bitset: %d is digit: %d\n", bs, dig)
		}

		for _, str := range strings.Split(output, " ") {
			for dig, bs := range keys {
				fmt.Printf("Testing %s (%d) against %d --> digit %d\n", str, encode(str), bs, dig)
				if bs == encode(str) {
					fmt.Printf("String %s (bitset %d) matches bitset %d, digit: %d\n", str, encode(str), bs, dig)
				}
			}
		}
	}
	fmt.Println(acc)
}

func init() {
	runeMap = map[rune]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
	}
	// Length: digit
	knownLen = map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}
}

func main() {
	ss = aoc.GetLines("shorttest.txt")
	p1()
	p2()
}

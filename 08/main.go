package main

import (
	aoc "aoc/utils"
	"fmt"
	"strings"
	"time"
)

var (
	ss     []string
	sigMap map[int]int
)

func p1() {
	defer aoc.TimeTrack(time.Now(), "Part one")
	var acc int
	for _, s := range ss {
		output := strings.Split(s, " | ")[1]
		for _, digit := range strings.Split(output, " ") {
			switch len(digit) {
			//   1, 4, 7, 8
			case 2, 4, 3, 7:
				acc++
			}
		}
	}
	fmt.Println(acc)
}

func strContainsAll(s, contained string) bool {
	for _, c := range contained {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}

func strCut(s, cutset string) (ret string) {
	for _, c := range s {
		if !strings.Contains(cutset, string(c)) {
			ret += string(c)
		}
	}
	return
}

func mapDigits(signal string) (digits [10]string) {
	toMap := []string{}
	for _, usig := range strings.Split(signal, " ") {
		sig := aoc.SortString(usig)
		switch len(sig) {
		case 2:
			digits[1] = sig
		case 3:
			digits[7] = sig
		case 4:
			digits[4] = sig
		case 7:
			digits[8] = sig
		default:
			toMap = append(toMap, sig)
		}
	}
	for _, sig := range toMap {
		if len(sig) == 6 {
			//if strings.Contains(sig, digits[4]) {
			if strContainsAll(sig, digits[4]) {
				digits[9] = sig
				continue
			}
			//if strings.Contains(sig, digits[1]) {
			if strContainsAll(sig, digits[1]) {
				digits[0] = sig
				continue
			}
			digits[6] = sig
			continue
		}
		if len(sig) == 5 {
			//if strings.Contains(sig, digits[1]) {
			if strContainsAll(sig, digits[1]) {
				digits[3] = sig
				continue
			}
			if len(strCut(sig, digits[4])) == 2 {
				digits[5] = sig
				continue
			}
			digits[2] = sig
		}
	}
	return
}

func decode(keys [10]string, signal string) (res int) {
	for _, unum := range strings.Split(signal, " ") {
		num := aoc.SortString(unum)
		for digit, k := range keys {
			if k == num {
				res = digit + (res * 10)
			}
		}
	}
	return
}

func p2() {
	defer aoc.TimeTrack(time.Now(), "Part two")
	var acc int
	for _, s := range ss {
		tmp := strings.Split(s, " | ")
		signals, output := tmp[0], tmp[1]
		digits := mapDigits(signals)
		res := decode(digits, output)
		acc += res
	}
	fmt.Println(acc)
}

func main() {
	ss = aoc.GetLines("input.txt")
	p1() // 445
	p2() // 1043101
}

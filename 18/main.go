package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"os"
	"time"
)

type tokenType int

type Token struct {
	typ   tokenType
	value int
}

type TokSlice []Token

const (
	OPEN tokenType = iota
	CLOSE
	COMMA
	NUMBER
)

var (
	allLines []string
	allToks  []TokSlice
)

func (ts TokSlice) String() string {
	var line string
	for _, tok := range ts {
		line += fmt.Sprint(tok)
	}
	return line
}

func (t Token) String() string {
	names := map[tokenType]string{
		OPEN: "[", CLOSE: "]", COMMA: ",",
	}
	if t.typ == NUMBER {
		return fmt.Sprintf("%d", t.value)
	}
	return fmt.Sprintf("%s", names[t.typ])
}

func (st TokSlice) Reduce() TokSlice {
	for exploded, split := true, true; exploded || split; {
		st, exploded = st.Explode()
		if exploded {
			continue
		}
		st, split = st.Split()
	}
	return st
}

func (st TokSlice) Explode() (newSlice TokSlice, change bool) {
	var open, openAt int
	for i := range st {
		if st[i].typ == OPEN {
			open++
			openAt = i
		}
		if st[i].typ == CLOSE && open > 4 {
			// Store values of current pair
			valLeft, valRight := st[openAt+1].value, st[i-1].value
			newSlice = make(TokSlice, len(st))
			copy(newSlice, st)

			// Search for number to the left
			for j := openAt - 1; j >= 0; j-- {
				if newSlice[j].typ == NUMBER {
					newSlice[j].value += valLeft
					break
				}
			}
			// Search for numbers to the right
			for j := i + 1; j < len(newSlice); j++ {
				if newSlice[j].typ == NUMBER {
					newSlice[j].value += valRight
					break
				}
			}

			// Replace the old pair with a 0 and assemble
			// the final slice
			newSlice = append(newSlice[:openAt],
				append([]Token{Token{typ: NUMBER}}, newSlice[i+1:]...)...)

			return newSlice, true
		}
		if st[i].typ == CLOSE {
			open--
		}
	}
	return st, false
}

func wrapPair(left, right Token) TokSlice {
	return TokSlice{
		Token{typ: OPEN}, left, Token{typ: COMMA}, right, Token{typ: CLOSE},
	}
}

func (st TokSlice) Split() (TokSlice, bool) {
	for i := range st {
		if st[i].typ == NUMBER {
			if st[i].value >= 10 {
				left := Token{NUMBER, st[i].value / 2}
				right := Token{NUMBER,
					int(math.Round(float64(st[i].value) / 2)),
				}
				return append(st[:i],
					append(wrapPair(left, right), st[i+1:]...)...), true
			}
		}
	}
	return st, false
}

func Add(st1 TokSlice, st2 TokSlice) TokSlice {
	newSlice := append([]Token{Token{typ: OPEN}}, st1...)
	newSlice = append(newSlice, Token{typ: COMMA})
	return append(newSlice, append(st2, Token{typ: CLOSE})...)
}

func Lex(line string) (toks TokSlice) {
	var num string
	tokTypes := map[rune]tokenType{
		'[': OPEN, ']': CLOSE, ',': COMMA,
	}
	for _, r := range line {
		if typ, in := tokTypes[r]; in {
			if num != "" {
				toks = append(toks, Token{
					typ:   NUMBER,
					value: aoc.ToInt(num),
				})
				num = ""
			}
			toks = append(toks, Token{typ: typ})
			continue
		}
		num += string(r)
	}
	return
}

func (ts TokSlice) Magnitude() int {
	for len(ts) > 1 {
		for i := 2; i < len(ts); i++ {
			if ts[i].typ == NUMBER && ts[i-2].typ == NUMBER {
				num := 3*ts[i-2].value + 2*ts[i].value
				ts = append(ts[:i-3], append([]Token{
					Token{typ: NUMBER, value: num}}, ts[i+2:]...)...)
			}
		}
	}
	return ts[0].value
}

func load(fname string) {
	aoc.TimeTrack(time.Now(), "Loading")
	if fname == "" {
		allLines = aoc.ReadStdin()
		fmt.Println(allLines)
		return
	}
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}
	allLines = aoc.GetLines(fname)
	for _, line := range allLines {
		allToks = append(allToks, Lex(line))
	}
}

func p1() {
	aoc.TimeTrack(time.Now(), "Part one")
	toks := make([]TokSlice, len(allToks))
	copy(toks, allToks)
	for len(toks) > 1 {
		toks[1] = Add(toks[0], toks[1]).Reduce()
		toks = toks[1:]
	}
	fmt.Println(toks[0].Magnitude())
}

func p2() {
	aoc.TimeTrack(time.Now(), "Part two")
	toks := make([]TokSlice, len(allToks))
	copy(toks, allToks)
	var max int
	for i := range toks {
		for j := range toks {
			if i == j {
				continue
			}
			magnitude := Add(toks[i], toks[j]).Reduce().Magnitude()
			if magnitude > max {
				max = magnitude
			}
		}
	}
	fmt.Println(max)
}

func main() {
	load("input.txt")
	p1()
	p2()
}

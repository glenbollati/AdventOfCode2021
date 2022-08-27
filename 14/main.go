package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

type insRules map[string]Insertion

type Insertion struct {
	char, triplet string
}

type Polymer struct {
	chain map[string]int64
	elems map[string]int64
}

var (
	rules    insRules
	template map[string]int64
	elements map[string]int64
)

func (p *Polymer) grow() {
	newChain := make(map[string]int64)
	for k, v := range p.chain {
		newChain[k] = v
	}
	for pair, ins := range rules {
		count := p.chain[pair]
		newChain[pair] -= count
		newChain[ins.triplet[:2]] += count
		newChain[ins.triplet[1:]] += count
		p.elems[ins.char] += count
	}
	p.chain = newChain
}

func (p *Polymer) elemCount() (most, least int64) {
	least = math.MaxInt64
	for _, count := range p.elems {
		if count < least {
			least = count
		}
		if count > most {
			most = count
		}
	}
	return
}

func newPolymer() *Polymer {
	p := Polymer{
		chain: map[string]int64{},
		elems: map[string]int64{},
	}
	for k, v := range template {
		p.chain[k] = v
	}
	for k, v := range elements {
		p.elems[k] = v
	}
	return &p
}

func load() {
	aoc.TimeTrack(time.Now(), "Loading")
	ss := aoc.GetSections("input.txt", "")

	// Template polymer
	template = make(map[string]int64)
	tmp := strings.Split(strings.Join(ss[0], ""), "")
	for i := range tmp[:len(tmp)-1] {
		template[strings.Join(tmp[i:i+2], "")]++
	}

	// Count elements
	elements = make(map[string]int64)
	for _, e := range tmp {
		elements[string(e)]++
	}

	// Insertion rules
	rules = make(insRules)
	for _, s := range ss[1] {
		spl := strings.Split(s, " -> ")
		old, new := spl[0], spl[1]
		rules[spl[0]] = Insertion{
			char:    new,
			triplet: string(old[0]) + new + string(old[1]),
		}
	}
}

func solve(steps int, label string) {
	defer aoc.TimeTrack(time.Now(), label)
	p := newPolymer()
	for i := 0; i < steps; i++ {
		p.grow()
	}
	most, least := p.elemCount()
	fmt.Println(most - least)
}

func main() {
	load()
	solve(10, "Part one")
	solve(40, "Part two")
}

package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"time"
)

type Pos [2]int

var (
	smallCave [][]int
	bigCave   [][]int
)

func (p *Pos) risk(cave [][]int) int {
	return cave[p[0]][p[1]]
}

type Pq struct {
	queue   map[int][]Pos
	minPrio int
}

func (p *Pq) Pop() Pos {
	if len(p.queue) == 0 {
		panic("Cannot pop from an empty queue")
	}
	mins := p.queue[p.minPrio]
	if len(mins) > 1 {
		pos := mins[0]
		p.queue[p.minPrio] = mins[1:]
		return pos
	}
	pos := mins[0]
	delete(p.queue, p.minPrio)
	p.minPrio = math.MaxInt
	for k, _ := range p.queue {
		if k < p.minPrio {
			p.minPrio = k
		}
	}
	return pos
}

func (p *Pq) Push(pos Pos, prio int) {
	if _, in := p.queue[prio]; in {
		p.queue[prio] = append(p.queue[prio], pos)
		return
	}
	p.queue[prio] = []Pos{pos}
	if prio < p.minPrio {
		p.minPrio = prio
	}
	return
}

func newPq() *Pq {
	return &Pq{
		queue:   make(map[int][]Pos),
		minPrio: math.MaxInt,
	}
}

// Dijkstra / A* - like solution
func findPath(cave [][]int) {
	limit := len(cave) - 1
	start := Pos{0, 0}
	end := Pos{limit, limit}

	toCheck := newPq()
	toCheck.Push(start, 0)

	riskToPos := map[Pos]int{start: 0}
	// parents := map[Pos]Pos{start: Pos{}}

	for len(toCheck.queue) > 0 {
		curr := toCheck.Pop()
		if curr == end {
			break
		}
		for _, adj := range adjacent(curr, limit) {
			newRisk := riskToPos[curr] + adj.risk(cave)
			if risk, in := riskToPos[adj]; !in || newRisk < risk {
				riskToPos[adj] = newRisk
				toCheck.Push(adj, newRisk)
				// parents[adj] = curr
			}
		}
	}
	fmt.Println(riskToPos[end])

	// If we needed the path this is how we would do it
	// path := []Pos{}
	// for pos := end; pos != start; pos = parents[pos] {
	// 	fmt.Println(pos, parents[pos])
	// 	path = append([]Pos{pos}, path...)
	// }
	// fmt.Println(path)
}

func adjacent(pos Pos, limit int) (adj []Pos) {
	row, col := pos[0], pos[1]
	if row < limit {
		adj = append(adj, Pos{row + 1, col})
	}
	if col < limit {
		adj = append(adj, Pos{row, col + 1})
	}
	if row > 0 {
		adj = append(adj, Pos{row - 1, col})
	}
	if col > 0 {
		adj = append(adj, Pos{row, col - 1})
	}
	return
}

func bumpRisk(risk int) int {
	risk += 1
	if risk > 9 {
		risk = risk - 9
	}
	return risk
}

// TODO: avoid building this whole thing
func makeBigCave() {
	bigCave = make([][]int, len(smallCave)*5)
	for row := range bigCave {
		bigCave[row] = make([]int, len(smallCave)*5)
	}

	for row := range smallCave {
		for col := range smallCave[row] {
			baseRisk := smallCave[row][col]
			risk := baseRisk
			// Walk across new row
			for i := 0; i < 5; i++ {
				bigCave[row][col+len(smallCave)*i] = risk
				risk = bumpRisk(risk)
			}
			risk = baseRisk
			// Walk down new columns
			for i := 0; i < 5; i++ {
				risk = baseRisk
				for b := 0; b < i; b++ {
					risk = bumpRisk(risk)
				}
				for j := 0; j < 5; j++ {
					bigCave[row+len(smallCave)*i][col+len(smallCave)*j] = risk
					risk = bumpRisk(risk)
				}
			}
		}
	}
}

func solve(cave [][]int, label string) {
	defer aoc.TimeTrack(time.Now(), label)
	findPath(cave)
}

func load() {
	defer aoc.TimeTrack(time.Now(), "Loading")
	smallCave = aoc.GetDigitGrid("input.txt")
	makeBigCave()
}

func main() {
	load()
	solve(smallCave, "Part one") // 687
	solve(bigCave, "Part two")   // 2957
}

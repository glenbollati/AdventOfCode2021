package main

import (
	aoc "aoc/utils"
	"fmt"
	"strings"
	"time"
)

var (
	caves map[string][]string
	paths int
)

type SearchState struct {
	usedTwice  bool
	visited    map[string]bool
	visitOrder []string
}

func makeState() SearchState {
	return SearchState{false, map[string]bool{}, []string{}}
}

// Copy over the state, adding the visited cave to it,
// Since maps and slices are pointers this takes care
// of making deep copies
func (s SearchState) visit(cave string) (ns SearchState) {
	ns = makeState()
	for vis := range s.visited {
		ns.visited[vis] = true
	}
	ns.visited[cave] = true
	ns.visitOrder = append(s.visitOrder, cave)
	ns.usedTwice = s.usedTwice
	return
}

// As with the visit() function above, but here we update the
// state to reflect the fact that we can no longer visit
// any small cave twice
func (s SearchState) useTwice() (ns SearchState) {
	ns = makeState()
	for vis := range s.visited {
		ns.visited[vis] = true
	}
	copy(ns.visitOrder, s.visitOrder)
	ns.usedTwice = true
	return
}

func dfs(start string, state SearchState, p1 bool) {
	state = state.visit(start)
	for _, next := range caves[start] {
		if next == "start" {
			continue
		}
		if next == "end" {
			paths++
			//fmt.Println(state.visit("end"))
			continue
		}
		if state.visited[next] && strings.ToLower(next) == next {
			if state.usedTwice || p1 {
				continue
			}
			ns := state.useTwice()
			dfs(next, ns, p1)
			continue
		}
		dfs(next, state, p1)
	}
	return
}

func solve(p1 bool, label string) {
	defer aoc.TimeTrack(time.Now(), label)
	paths = 0
	dfs("start", makeState(), p1)
	fmt.Println(paths)
}

func main() {
	ss := aoc.GetLines("input.txt")
	caves = make(map[string][]string)
	for _, s := range ss {
		spl := strings.Split(s, "-")
		id, adj := spl[0], spl[1]
		caves[id] = append(caves[id], adj)
		caves[adj] = append(caves[adj], id)
	}
	solve(true, "Part one")  // 4754
	solve(false, "Part two") // 143562
}

/*
func (s SearchState) String() string {
	var str string
	for _, cave := range s.visitOrder {
		str += cave + ","
	}
	return strings.TrimSuffix(str, ",")
}

func isSmall(cave string) bool {
	for _, r := range cave {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
*/

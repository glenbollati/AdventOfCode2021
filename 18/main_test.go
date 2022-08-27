package main

import (
	"fmt"
	"testing"
)

func TestExplode(t *testing.T) {
	tests := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}
	for input, want := range tests {
		got, _ := Lex(input).Explode()
		if fmt.Sprint(got) != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := map[string]string{
		//"10":                              "[5,5]",
		//"11":                              "[5,6]",
		//"12":                              "[6,6]",
		"[[[[0,7],4],[15,[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
	}
	for input, want := range tests {
		got, _ := Lex(input).Split()
		if fmt.Sprint(got) != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := map[string][]string{
		"[[1,2],[[3,4],5]]":                     []string{"[1,2]", "[[3,4],5]"},
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]": []string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"},
	}
	for want, input := range tests {
		got := Add(Lex(input[0]), Lex(input[1]))
		if fmt.Sprint(got) != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	}
}

/*(
func TestReduce(t *testing.T) {
	tests := map[string]string{
		"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]": "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
	}
	for input, want := range tests {
		got := Lex(input).Reduce()
		if fmt.Sprint(got) != want {
			t.Errorf("Got:\n%s\nWanted:\n%s\n", got, want)
		}
	}
}
*/

func TestMagnitude(t *testing.T) {
	tests := map[string]int{
		"[9,1]":                             29,
		"[[9,1],[1,9]]":                     129,
		"[[1,2],[[3,4],5]]":                 143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]": 1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":     445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":     791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":     1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}
	for input, want := range tests {
		got := Lex(input).Magnitude()
		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	}
}

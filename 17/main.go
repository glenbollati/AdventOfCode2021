package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"os"
)

type i2 struct{ x, y int }

type Probe struct {
	curr, prev, vel i2
}

var (
	target struct{ min, max i2 }
)

func (p *Probe) inTarget() bool {
	if p.curr.x <= target.max.x && p.curr.x >= target.min.x &&
		p.curr.y <= target.max.y && p.curr.y >= target.min.y {
		return true
	}
	return false
}

func (a *i2) dist(b i2) int {
	dx, dy := b.x-a.x, b.y-a.y
	return int(math.Round(math.Sqrt(float64((dx * dx) + (dy * dy)))))
}

func load(fname string) {
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}
	fmt.Sscanf(aoc.ReadFile(fname), "target area: x=%d..%d, y=%d..%d",
		&target.min.x, &target.max.x, &target.min.y, &target.max.y)
}

func simulate(initVel i2) (success bool, ymax int) {
	p := Probe{vel: initVel}
	for {
		// DEBUG
		//	fmt.Println(p.curr, target)
		if p.inTarget() {
			return true, ymax
		}
		// Find if we overshot: passed through target during last step
		if p.prev.y > target.max.y && p.curr.y < target.min.y {
			return false, ymax
		}
		if p.prev.x < target.min.x && p.curr.x > target.max.x {
			return false, ymax
		}
		if p.curr.x == p.prev.x {
			if p.curr.y < target.min.y {
				// DEBUG fmt.Println("bailing on y < miny")
				return false, ymax
			}
		}
		p.prev = p.curr
		p.curr.x += p.vel.x
		p.curr.y += p.vel.y
		p.vel.y--
		// Keep track of our max y position
		if p.curr.y > ymax {
			ymax = p.curr.y
		}

		if p.vel.x > 0 {
			p.vel.x--
			continue
		}
		if p.vel.x < 0 {
			p.vel.x++
		}
	}
	panic("Unreachable")
}

// How high can you make the probe go while still reaching the target area?
// Find the initial velocity that causes the probe to reach the highest
// y position and still eventually be within the target area after any step.
// What is the highest y position it reaches on this trajectory?
func p1() {
	var min i2
	x := aoc.Min(0, target.min.x)
	for {
		// fmt.Println("Simulating with x =", x)
		if success, _ := simulate(i2{x, 0}); success {
			// fmt.Printf("x=%d ymax=%d\n", x, ymax)
			break
		}
		x++
	}
	min.x = x

	// Start with y at 0, increase until we hit success
	y := aoc.Min(0, target.min.y)
	for {
		// fmt.Println("Simulating with y =", y)
		if success, _ := simulate(i2{min.x, y}); success {
			// fmt.Printf("Found ymin: %d\n", y)
			break
		}
		y++
	}
	min.y = y

	// Increase y from ymin until we no longer hit success
	y, ymax, yMAX := 0, 0, 0
	success := false
	for {
		// fmt.Printf("Trying with y = %d\n", y)
		if success, ymax = simulate(i2{min.x, y}); !success {
			// fmt.Printf("Found limit: %d\n", y)
			break
		}
		yMAX = aoc.Max(yMAX, ymax)
		y++
	}
	fmt.Println(yMAX)

	//for _, initVel := range []i2{i2{6, 2}, i2{7, 2}, i2{6, 3}, i2{7, 3}} {
	//	res, ymax := simulate(initVel)
	//	fmt.Println(initVel, res, ymax)
	//}
	for i := -1000; i < 1000; i++ {
		for j := -1000; j < 1000; j++ {
			initVel := i2{i, j}
			res, ymax := simulate(initVel)
			if res {
				fmt.Println(initVel, res, ymax)
			}
		}
	}
	for i := -1000; i < 1000; i++ {
		for j := -1000; j < 1000; j++ {
			initVel := i2{j, i}
			res, ymax := simulate(initVel)
			if res {
				fmt.Println(initVel, res, ymax)
			}
		}
	}
	//for _, initVel := range []i2{i2{22, 40}, i2{22, 44}} {
	//	res, ymax := simulate(initVel)
	//	fmt.Println(initVel, res, ymax)
	//}
}

func main() {
	load("input.txt")
	p1() // {21 88} true 3916
	//p2() // 2986
}

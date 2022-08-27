package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"strings"
)

/*
	In total, each scanner could be in any of 24 different orientations:
	facing positive or negative x, y, or z,
	and considering any of four directions "up" from that facing.

	Given a scanner and a map of its relative beacon positions, find a way to
	place other scanners within that map such that at least 12 beacons overlap

	Distance between beacons is fixed
	Use a beacon as fixed point, find distances from the beacon to other beacons
	Do the same for beacon from the other set
	map beacons from each set to each other based on the distance between them
*/

type distMap map[int][2]i3

type i3 struct {
	x, y, z int
}

var scanData [][]i3

func (a i3) Sub(b i3) i3 {
	return i3{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}

// NOTE: may have to use float as result
func (a i3) Dist(b i3) int {
	return int(math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2) + math.Pow(float64(b.z-a.z), 2)))
}

func load(fname string) {
	ss := aoc.GetSections(fname, "")
	scanData = make([][]i3, len(ss))
	dataPoint := i3{}
	for scanner, data := range ss {
		for _, line := range data {
			if strings.Contains(line, "scanner") {
				continue
			}
			fmt.Sscanf(line, "%d,%d,%d", &dataPoint.x, &dataPoint.y, &dataPoint.z)
			scanData[scanner] = append(scanData[scanner], dataPoint)
		}
	}
}

// Given a scanner number, mapBeacons builds a map of distances
// between the beacons detected by the scanner
func mapBeacons(scanner int) distMap {
	baseData := scanData[scanner]
	dm, mapped := make(distMap), make(map[[2]i3]bool)
	for i, bi := range baseData {
		for j, bj := range baseData {
			//fmt.Printf("Checking %+v against %+v, distance: %d\n", bi, bj, bi.Dist(bj))
			combo := [2]i3{bi, bj}
			if i == j || mapped[combo] {
				continue
			}
			mapped[combo] = true
			dist := bi.Dist(bj)
			dm[dist] = combo
		}
	}
	return dm
}

func p1() {
	uniqueDists := make(map[int]bool)
	count := 0
	for scanner := range scanData[:2] {
		for dist := range mapBeacons(scanner) {
			if uniqueDists[dist] {
				count++
				continue
			}
			uniqueDists[dist] = true
		}
	}
	fmt.Println(count)
}

func main() {
	load("test.txt")
	p1()
}

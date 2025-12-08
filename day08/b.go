package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x, y, z int
	circuit int
}

type Pair struct {
	distance float64
	i, j     int
}

var circuits = make(map[int][]int)

func readJunctionBoxes() []JunctionBox {
	var boxes []JunctionBox
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		boxId := len(boxes)
		boxes = append(boxes, JunctionBox{x, y, z, boxId})
		circuits[boxId] = []int{boxId}
	}
	return boxes
}

func computeDistances(boxes []JunctionBox) []Pair {
	n := len(boxes)
	var distances []Pair
	for i := range boxes {
		for j := i + 1; j < n; j++ {
			xd := float64(boxes[i].x - boxes[j].x)
			yd := float64(boxes[i].y - boxes[j].y)
			zd := float64(boxes[i].z - boxes[j].z)
			distance := math.Sqrt(xd*xd + yd*yd + zd*zd)
			distances = append(distances, Pair{distance, i, j})
		}
	}

	slices.SortFunc(distances, func(a, b Pair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return distances
}

func union(boxes []JunctionBox, from, to int) {
	for _, box := range circuits[from] {
		boxes[box].circuit = to
	}
	circuits[to] = append(circuits[to], circuits[from]...)
	delete(circuits, from)
}

func main() {
	junctionBoxes := readJunctionBoxes()
	distances := computeDistances(junctionBoxes)

	di := 0
	for {
		pair := distances[di]
		di++
		boxA := junctionBoxes[pair.i]
		boxB := junctionBoxes[pair.j]
		if boxA.circuit == boxB.circuit {
			// On the same circuit already
			continue
		}
		if len(circuits[boxA.circuit]) > len(circuits[boxB.circuit]) {
			union(junctionBoxes, boxB.circuit, boxA.circuit)
		} else {
			union(junctionBoxes, boxA.circuit, boxB.circuit)
		}
		if len(circuits) == 1 {
			product := boxA.x * boxB.x
			fmt.Println(product)
			break
		}
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var timelines int
var cache map[string]int
var grid [][]byte

func readGrid() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
}

func quantumSplit(row int, beam int) int {
	if row >= len(grid)-1 {
		return 1
	}
	if grid[row][beam] == byte('^') {
		d := 0
		if beam-1 >= 0 {
			key := fmt.Sprintf("%d-%d", row+1, beam-1)
			left, ok := cache[key]
			if !ok {
				left = quantumSplit(row+1, beam-1)
				cache[key] = left
			}
			d += left
		}
		if beam+1 < len(grid[row]) {
			key := fmt.Sprintf("%d-%d", row+1, beam+1)
			right, ok := cache[key]
			if !ok {
				right = quantumSplit(row+1, beam+1)
				cache[key] = right
			}
			d += right
		}
		return d
	}
	key := fmt.Sprintf("%d-%d", row+1, beam)
	mid, ok := cache[key]
	if ok {
		return mid
	}
	mid = quantumSplit(row+1, beam)
	cache[key] = mid
	return mid
}

func simulate() {
	cache = make(map[string]int)
	start := bytes.Index(grid[0], []byte{'S'})
	timelines = quantumSplit(1, start)
}

func main() {
	readGrid()
	simulate()
	fmt.Println(timelines)
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func readGrid() [][]byte {
	var grid [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	return grid
}

var timelines int
var cache map[string]int

func quantumSplit(grid [][]byte, row int, path string, beam int) int {
	if row >= len(grid)-1 {
		return 1
	}
	if grid[row][beam] == byte('^') {
		d := 0
		if beam-1 >= 0 {
			key := fmt.Sprintf("%d-%d", row+1, beam-1)
			left, ok := cache[key]
			if !ok {
				left = quantumSplit(grid, row+1, path+"L", beam-1)
				cache[key] = left
			}
			d += left
		}
		if beam+1 < len(grid[row]) {
			key := fmt.Sprintf("%d-%d", row+1, beam+1)
			right, ok := cache[key]
			if !ok {
				right = quantumSplit(grid, row+1, path+"R", beam+1)
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
	mid = quantumSplit(grid, row+1, path, beam)
	cache[key] = mid
	return mid
}

func simulate(grid [][]byte) {
	cache = make(map[string]int)
	start := bytes.Index(grid[0], []byte{'S'})
	timelines = quantumSplit(grid, 1, "", start)
}

func main() {
	grid := readGrid()
	simulate(grid)
	fmt.Println(timelines)
}

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

func simulate(grid [][]byte) int {
	start := bytes.Index(grid[0], []byte{'S'})
	splitCount := 0
	beams := map[int]int{start: 1}
	for row := 1; row < len(grid); row++ {
		for k, _ := range beams {
			if grid[row][k] == byte('^') {
				splitCount++
				if k-1 >= 0 {
					beams[k-1] = 1
				}
				if k+1 < len(grid[row]) {
					beams[k+1] = 1
				}
				delete(beams, k)
			}
		}
	}
	return splitCount
}

func main() {
	grid := readGrid()
	splitCount := simulate(grid)
	fmt.Println(splitCount)
}

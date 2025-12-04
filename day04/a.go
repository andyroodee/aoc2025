package main

import (
	"bufio"
	"fmt"
	"os"
)

func readGrid() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func inBounds(grid []string, row, col int) bool {
	return row >= 0 && row < len(grid) &&
		col >= 0 && col < len(grid[row])
}

func accessible(grid []string, row, col int) int {
	dirs := []int{-1, 0, 1}
	count := 0
	for _, y := range dirs {
		for _, x := range dirs {
			r := row + y
			c := col + x
			if (r == row && c == col) || !inBounds(grid, r, c) {
				continue
			}
			if grid[r][c] == '@' {
				count++
			}
		}
	}
	if count < 4 {
		return 1
	}
	return 0
}

func main() {
	grid := readGrid()
	rolls := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '@' {
				rolls += accessible(grid, row, col)
			}
		}
	}

	fmt.Println(rolls)
}

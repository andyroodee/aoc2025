package main

import (
	"bufio"
	"fmt"
	"os"
)

func readGrid() [][]byte {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	return grid
}

func inBounds(grid [][]byte, row, col int) bool {
	return row >= 0 && row < len(grid) &&
		col >= 0 && col < len(grid[row])
}

func accessible(grid [][]byte, row, col int) bool {
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
	return count < 4
}

func main() {
	grid := readGrid()
	totalRolls := 0
	for {
		rolls := 0
		nextGrid := make([][]byte, len(grid))
		for i, row := range grid {
			nextGrid[i] = make([]byte, len(row))
			copy(nextGrid[i], row)
		}
		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] == '@' {
					if accessible(grid, row, col) {
						rolls++
						nextGrid[row][col] = '.'
					}
				}
			}
		}
		if rolls == 0 {
			break
		}
		totalRolls += rolls
		grid = nextGrid
	}

	fmt.Println(totalRolls)
}

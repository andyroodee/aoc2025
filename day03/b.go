package main

import (
	"bufio"
	"fmt"
	"os"
)

func readBanks() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	var banks [][]int
	for scanner.Scan() {
		line := scanner.Text()
		batteries := make([]int, len(line))
		for i, b := range line {
			batteries[i] = int(b - '0')
		}
		banks = append(banks, batteries)
	}
	return banks
}

func getMaxJoltage(bank []int) int {
	const goal = 12
	best := make([]int, len(bank)+1)
	mult := 1
	for size := 1; size < goal+1; size++ {
		next := make([]int, len(best)-1)
		for i := len(bank) - size; i >= 0; i-- {
			n := mult*bank[i] + best[i+1]
			if i == len(bank)-size {
				next[i] = n
			} else {
				if n > next[i+1] {
					next[i] = n
				} else {
					next[i] = next[i+1]
				}
			}
		}
		best = next
		mult *= 10
	}
	return best[0]
}

func main() {
	banks := readBanks()
	var totalJoltage int
	for _, bank := range banks {
		totalJoltage += getMaxJoltage(bank)
	}
	fmt.Println(totalJoltage)
}

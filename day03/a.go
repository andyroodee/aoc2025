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
	var maxJoltage int
	maxLeft := bank[0]
	for i := 0; i < len(bank)-1; i++ {
		left := bank[i]
		if left < maxLeft {
			continue
		}
		for j := i + 1; j < len(bank); j++ {
			right := bank[j]
			sum := 10*left + right
			if sum > maxJoltage {
				maxJoltage = sum
			}
			if right == 9 {
				break
			}
		}
	}
	return maxJoltage
}

func main() {
	banks := readBanks()
	var totalJoltage int
	for _, bank := range banks {
		totalJoltage += getMaxJoltage(bank)
	}
	fmt.Println(totalJoltage)
}

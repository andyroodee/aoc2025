package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dial := 50
	zeroCount := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		amount, _ := strconv.Atoi(line[1:])
		fullTurns := amount / 100
		zeroCount += fullTurns
		before := dial
		if dir == 'L' {
			dial -= amount % 100
		} else {
			dial += amount % 100
		}
		if dial < 0 {
			dial += 100
			if before > 0 {
				zeroCount++
			}
		}
		if dial > 99 {
			dial -= 100
			if dial > 0 {
				zeroCount++
			}
		}
		if dial == 0 {
			zeroCount++
		}
	}
	fmt.Println(zeroCount)
}

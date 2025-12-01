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
		//fmt.Printf("dir: %c, amount: %d\n", dir, amount)
		if dir == 'L' {
			dial -= amount % 100
		} else {
			dial += amount % 100
		}
		if dial < 0 {
			dial += 100
		}
		if dial > 99 {
			dial -= 100
		}
		if dial == 0 {
			zeroCount++
		}
	}
	fmt.Println(zeroCount)
}

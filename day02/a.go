package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidId(s string) bool {
	if len(s)%2 != 0 {
		return true
	}
	mid := len(s) / 2
	return s[:mid] != s[mid:]
}

func main() {
	var invalidIdSum int
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	ranges := strings.Split(strings.TrimSpace(line), ",")
	for _, r := range ranges {
		vals := strings.Split(r, "-")
		lo, _ := strconv.Atoi(vals[0])
		hi, _ := strconv.Atoi(vals[1])
		for i := lo; i <= hi; i++ {
			iStr := strconv.Itoa(i)
			if !isValidId(iStr) {
				invalidIdSum += i
			}
		}
	}
	fmt.Println(invalidIdSum)
}

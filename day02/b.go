package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidId(s string) bool {
	length := len(s)
	if length < 2 {
		return true
	}
	mid := length / 2
	for i := 1; i <= mid; i++ {
		if length%i != 0 {
			continue
		}
		invalid := true
		comp := s[:i]
		for j := i; j <= length-i; j += i {
			other := s[j : j+i]
			if comp != other {
				invalid = false
				break
			}
		}
		if invalid {
			return false
		}
	}
	return true
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

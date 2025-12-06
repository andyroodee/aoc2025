package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	numbers []int
	op      string
}

type operator func(a, b int) int

var ops = map[string]operator{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

func readProblems() []problem {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	var problems []problem
	for scanner.Scan() {
		line := scanner.Text()
		entries := strings.Fields(line)
		if count == 0 {
			count = len(entries)
			problems = make([]problem, count)
		}
		if entries[0] == "*" || entries[0] == "+" {
			for i := range entries {
				problems[i].op = entries[i]
			}
		} else {
			for i := range entries {
				num, _ := strconv.Atoi(entries[i])
				problems[i].numbers = append(problems[i].numbers, num)
			}
		}
	}
	return problems
}

func calculate(problems []problem) int {
	grandTotal := 0
	for _, p := range problems {
		op := ops[p.op]
		total := p.numbers[0]
		for i := 1; i < len(p.numbers); i++ {
			total = op(total, p.numbers[i])
		}
		grandTotal += total
	}
	return grandTotal
}

func main() {
	problems := readProblems()
	total := calculate(problems)
	fmt.Println(total)
}

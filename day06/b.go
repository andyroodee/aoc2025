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

func readColumn(col int, grid []string) string {
	var sb strings.Builder
	for row := 0; row < len(grid); row++ {
		sb.WriteByte(grid[row][col])
	}
	return strings.TrimSpace(sb.String())
}

func readProblems() []problem {
	scanner := bufio.NewScanner(os.Stdin)
	var problems []problem
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	cols := len(grid[0])
	for col := 0; col < cols; col++ {
		colStr := readColumn(col, grid)
		if colStr == "" {
			continue
		}
		last := colStr[len(colStr)-1]
		if last == '+' || last == '*' {
			problem := problem{
				op: string(last),
			}
			n, _ := strconv.Atoi(strings.TrimSpace(colStr[:len(colStr)-1]))
			problem.numbers = append(problem.numbers, n)
			problems = append(problems, problem)
		} else {
			n, _ := strconv.Atoi(colStr)
			problems[len(problems)-1].numbers = append(problems[len(problems)-1].numbers, n)
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

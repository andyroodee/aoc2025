package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type graph map[string][]string

var g = graph{}

const goal = "out"

func readGraph() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		key := fields[0]
		key = key[:len(key)-1]
		for _, link := range fields[1:] {
			g[key] = append(g[key], link)
		}
	}
}

func solve(start string) int {
	if start == goal {
		return 1
	}
	sum := 0
	for _, link := range g[start] {
		sum += solve(link)
	}
	return sum
}

func main() {
	readGraph()
	pathCount := solve("you")
	fmt.Println(pathCount)
}

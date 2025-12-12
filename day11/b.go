package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type graph map[string][]string

var g = graph{}
var memo = map[string]int{}

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

func solve(start string, dac, fft bool) int {
	if start == goal {
		if dac && fft {
			return 1
		}
		return 0
	}
	key := fmt.Sprintf("%s%v%v", start, dac, fft)
	s, ok := memo[key]
	if ok {
		return s
	}
	if !dac && slices.Contains(g["dac"], start) {
		memo[key] = 0
		return 0
	}
	if !fft && slices.Contains(g["fft"], start) {
		memo[key] = 0
		return 0
	}
	sum := 0
	dac = dac || start == "dac"
	fft = fft || start == "fft"
	for _, link := range g[start] {
		sum += solve(link, dac, fft)
	}
	memo[key] = sum
	return sum
}

func main() {
	readGraph()
	pathCount := solve("svr", false, false)
	fmt.Println(pathCount)
}

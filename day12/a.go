package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gift []string

type area struct {
	width    int
	length   int
	presents []int
}

func nextLine(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func readInput() ([]gift, []area) {
	var gifts []gift
	var areas []area
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		x := strings.Index(line, "x")
		if x == -1 {
			g := gift{}
			g = append(g, nextLine(scanner))
			g = append(g, nextLine(scanner))
			g = append(g, nextLine(scanner))
			gifts = append(gifts, g)
			nextLine(scanner)
		} else {
			a := area{}
			parts := strings.Fields(line)
			dim := parts[0][:len(parts[0])-1]
			w, _ := strconv.Atoi(dim[:x])
			l, _ := strconv.Atoi(dim[x+1:])
			a.width = w
			a.length = l
			for _, p := range parts[1:] {
				index, _ := strconv.Atoi(p)
				a.presents = append(a.presents, index)
			}
			areas = append(areas, a)
		}
	}
	return gifts, areas
}

func main() {
	_, areas := readInput()
	const giftArea = 9
	count := 0
	for _, a := range areas {
		size := a.width * a.length
		req := 0
		for _, p := range a.presents {
			req += p * giftArea
		}
		if req <= size {
			count++
		}
	}
	fmt.Println(count)
}

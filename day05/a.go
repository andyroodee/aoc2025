package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	lo, hi int
}

func readRanges(scanner *bufio.Scanner) []idRange {
	var ranges []idRange
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, "-")
		lo, _ := strconv.Atoi(split[0])
		hi, _ := strconv.Atoi(split[1])
		r := idRange{
			lo: lo,
			hi: hi,
		}
		ranges = append(ranges, r)
	}
	return ranges
}

func readIngredients(scanner *bufio.Scanner) []int {
	var ingredients []int
	for scanner.Scan() {
		line := scanner.Text()
		ingredient, _ := strconv.Atoi(line)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ranges := readRanges(scanner)
	ingredients := readIngredients(scanner)
	count := 0
	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r.lo && ingredient <= r.hi {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type idRange struct {
	lo, hi int
}

func (r *idRange) overlapsWith(other idRange) bool {
	return r.lo <= other.hi && other.lo <= r.hi
}

func (r *idRange) mergeWith(other idRange) {
	if other.lo < r.lo {
		r.lo = other.lo
	}
	if other.hi > r.hi {
		r.hi = other.hi
	}
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
	slices.SortFunc(ranges, func(a, b idRange) int {
		return cmp.Compare(a.lo, b.lo)
	})
	return ranges
}

func mergeRanges(ranges []idRange) ([]idRange, bool) {
	var merged []idRange
	hasMerged := false
	for i := 0; i < len(ranges); i++ {
		a := ranges[i]
		if i+1 < len(ranges) && a.overlapsWith(ranges[i+1]) {
			a.mergeWith(ranges[i+1])
			i++
			hasMerged = true
		}
		merged = append(merged, a)
	}
	return merged, hasMerged
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ranges := readRanges(scanner)
	for {
		mergedRanges, hasMerged := mergeRanges(ranges)
		if !hasMerged {
			break
		}
		ranges = mergedRanges
	}
	count := 0
	for _, r := range ranges {
		count += r.hi - r.lo + 1
	}
	fmt.Println(count)
}

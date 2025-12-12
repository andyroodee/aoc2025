package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type button []int

type constraint struct {
	buttons []int
	signs   []int
	value   int
}

type machine struct {
	buttons     []button
	buttonOrder []int
	goaltages   []int
	pressLimits []int
	constraints []constraint
	smallest    int
}

func search(m *machine, k int, n int, counts []int) {
	if k == n {
		totals := make([]int, len(m.goaltages))
		sum := 0
		for i, c := range counts {
			for _, b := range m.buttons[i] {
				totals[b] += c
			}
			sum += c
		}
		if reflect.DeepEqual(totals, m.goaltages) {
			if sum < m.smallest {
				m.smallest = sum
			}
		}
		return
	}
	bi := m.buttonOrder[k]
	for i := range m.pressLimits[bi] + 1 {
		counts[bi] = i
		ok := true
		for _, c := range m.constraints {
			// k is the button index. buttons <= k are 'set' for this iteration and can be checked for constraints.
			if slices.Contains(c.buttons, bi) && !slices.ContainsFunc(c.buttons, func(x int) bool { return slices.Index(m.buttonOrder, x) > k }) {
				sum := 0
				for ind, b := range c.buttons {
					sum += counts[b] * c.signs[ind]
				}
				if sum != c.value {
					ok = false
					break
				}
			}
		}
		if ok {
			search(m, k+1, n, counts)
		}
	}
}

func makeMachine(buttonDefs []string, joltageDef string) machine {
	m := machine{
		smallest: math.MaxInt,
	}

	jolts := strings.Split(joltageDef[1:len(joltageDef)-1], ",")
	for _, j := range jolts {
		val, _ := strconv.Atoi(j)
		m.goaltages = append(m.goaltages, val)
	}

	for _, bd := range buttonDefs {
		var bs button
		buttonGroup := strings.Split(bd[1:len(bd)-1], ",")
		maxPress := math.MaxInt
		for _, bg := range buttonGroup {
			val, _ := strconv.Atoi(bg)
			bs = append(bs, val)
			if m.goaltages[val] < maxPress {
				maxPress = m.goaltages[val]
			}
		}
		m.buttons = append(m.buttons, bs)
		m.pressLimits = append(m.pressLimits, maxPress)
	}

	for i, g := range m.goaltages {
		c := constraint{
			value: g,
		}
		for j, b := range m.buttons {
			if slices.Contains(b, i) {
				c.buttons = append(c.buttons, j)
				c.signs = append(c.signs, 1)
			}
		}
		if !slices.ContainsFunc(m.constraints, func(x constraint) bool {
			return constraintsEqual(x, c)
		}) {
			m.constraints = append(m.constraints, c)
		}
	}

	// Extra constraints
	var extra []constraint
	for i := 0; i < len(m.constraints); i++ {
		a := m.constraints[i]
		for j := i + 1; j < len(m.constraints); j++ {
			b := m.constraints[j]
			if a.value > b.value {
				bonus, ok := checkForExtraConstraint(b, a, extra, m.constraints)
				if ok {
					extra = append(extra, bonus)
				}
			} else {
				bonus, ok := checkForExtraConstraint(a, b, extra, m.constraints)
				if ok {
					extra = append(extra, bonus)
				}
			}
		}
	}
	if len(extra) > 0 {
		m.constraints = append(m.constraints, extra...)
	}

	slices.SortFunc(m.constraints, func(a, b constraint) int {
		return cmp.Compare(len(a.buttons), len(b.buttons))
	})
	buttonMap := make(map[int]bool)
	for _, c := range m.constraints {
		for _, b := range c.buttons {
			if !buttonMap[b] {
				buttonMap[b] = true
				m.buttonOrder = append(m.buttonOrder, b)
			}
		}
	}

	return m
}

func checkForExtraConstraint(small, big constraint, extra []constraint, m []constraint) (constraint, bool) {
	ok := false
	for _, test := range small.buttons {
		if slices.Contains(big.buttons, test) {
			ok = true
			break
		}
	}
	if !ok {
		return constraint{}, false
	}
	bonus := constraint{}
	for _, test := range big.buttons {
		if !slices.Contains(small.buttons, test) {
			bonus.buttons = append(bonus.buttons, test)
			bonus.signs = append(bonus.signs, 1)
		}
	}
	for _, test := range small.buttons {
		if !slices.Contains(big.buttons, test) {
			bonus.buttons = append(bonus.buttons, test)
			bonus.signs = append(bonus.signs, -1)
		}
	}
	bonus.value = big.value - small.value
	bonusLen := len(bonus.buttons)

	if bonusLen > 0 && bonusLen < len(small.buttons) && bonusLen < len(big.buttons) {
		if !slices.ContainsFunc(extra, func(x constraint) bool {
			return constraintsEqual(x, bonus)
		}) && !slices.ContainsFunc(m, func(x constraint) bool {
			return constraintsEqual(x, bonus)
		}) {
			return bonus, true
		}
	}
	return constraint{}, false
}

func constraintsEqual(a, b constraint) bool {
	if a.value != b.value {
		return false
	}
	if len(a.buttons) != len(b.buttons) {
		return false
	}
	for i := range a.buttons {
		u := a.buttons[i] * a.signs[i]
		v := b.buttons[i] * b.signs[i]
		if u != v {
			return false
		}
	}
	return true
}

func readMachines() []machine {
	var machines []machine
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		buttonDefs := parts[1 : len(parts)-1]
		joltageDef := parts[len(parts)-1]
		m := makeMachine(buttonDefs, joltageDef)
		machines = append(machines, m)
	}
	return machines
}

func main() {
	minSum := 0
	machines := readMachines()
	var wg sync.WaitGroup
	sums := make([]int, len(machines))
	for i, m := range machines {
		wg.Go(func() {
			search(&m, 0, len(m.buttons), make([]int, len(m.buttons)))
			fmt.Printf("machine %d min: %d\n", i, m.smallest)
			sums[i] = m.smallest
		})
	}
	wg.Wait()
	for _, sum := range sums {
		minSum += sum
	}
	fmt.Println(minSum)
}

func printConstraints(constraints []constraint) {
	for _, c := range constraints {
		var sb strings.Builder
		for i, b := range c.buttons {
			p := byte(b) + 'a'
			if i == 0 && c.signs[i] == -1 {
				sb.WriteByte('-')
			}
			if i > 0 {
				if c.signs[i] == -1 {
					sb.WriteString(" - ")
				} else {
					sb.WriteString(" + ")
				}
			}
			sb.WriteByte(p)
		}
		fmt.Println(sb.String(), "=", c.value)
	}
}

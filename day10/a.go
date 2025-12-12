package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type toggle uint16

type search struct {
	lights   toggle
	distance int
}

type machine struct {
	goalState toggle
	buttons   []toggle
}

func newMachine(lightsStr string, buttonsStr []string) machine {
	var goalState toggle
	lightCount := len(lightsStr)
	for i, x := range lightsStr {
		bit := lightCount - 1 - i
		if x == '#' {
			goalState |= 1 << bit
		}
	}

	var buttons []toggle
	for _, buttonStr := range buttonsStr {
		buttons = append(buttons, makeButton(lightCount, buttonStr))
	}

	return machine{
		goalState: goalState,
		buttons:   buttons,
	}
}

func makeButton(lightCount int, buttonStr string) toggle {
	var button toggle
	parts := strings.Split(buttonStr[1:len(buttonStr)-1], ",")
	for _, part := range parts {
		b, _ := strconv.Atoi(part)
		bit := lightCount - 1 - b
		button |= 1 << bit
	}
	return button
}

func readInput() []machine {
	var machines []machine
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		lights := fields[0]
		lights = lights[1 : len(lights)-1]
		buttons := fields[1 : len(fields)-1]
		// Ignore joltage requirements for now
		//joltages := fields[len(fields)-1]
		machines = append(machines, newMachine(lights, buttons))
	}
	return machines
}

func solve(machines []machine) {
	sum := 0
	for _, m := range machines {
		seen := make(map[toggle]bool)
		var startState toggle
		q := []search{{startState, 0}}
		for len(q) > 0 {
			next := q[0]
			q = q[1:]
			if next.lights == m.goalState {
				sum += next.distance
				break
			}
			for _, button := range m.buttons {
				if seen[next.lights^button] {
					// If we've already seen this configuration there's no shorter path
					continue
				}
				q = append(q, search{next.lights ^ button, next.distance + 1})
			}
			seen[next.lights] = true
		}
	}
	fmt.Println(sum)
}

func main() {
	machines := readInput()
	solve(machines)
}

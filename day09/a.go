package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func readPoints() []Point {
	scanner := bufio.NewScanner(os.Stdin)
	var points []Point
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		point := Point{
			x: x,
			y: y,
		}
		points = append(points, point)
	}
	return points
}

func getMaxRectangleArea(points []Point) int {
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := int(math.Abs(float64(a.x-b.x)+1) * math.Abs(float64(a.y-b.y)+1))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {
	points := readPoints()
	maxArea := getMaxRectangleArea(points)
	fmt.Println(maxArea)
}

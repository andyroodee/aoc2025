package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

const puzzleSize = 100000

type State byte

const (
	Empty State = iota
	Wall
	Out
	Seen
)

var points []Point
var gridState = make([][]State, puzzleSize)

func readPoints() {
	scanner := bufio.NewScanner(os.Stdin)
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
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Rect struct {
	minX, minY, maxX, maxY int
}

func (r Rect) contains(other Rect) bool {
	return r.minX <= other.minX && r.minY <= other.minY && other.maxX <= r.maxX && other.maxY <= r.maxY
}

var badRects = make(map[Rect]bool)

func getMaxRectangleArea() int {
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			minX := intMin(a.x, b.x)
			maxX := intMax(a.x, b.x)
			minY := intMin(a.y, b.y)
			maxY := intMax(a.y, b.y)
			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area <= maxArea {
				continue
			}
			inBounds := true
			rect := Rect{minX, minY, maxX, maxY}
			for k := range badRects {
				if rect.contains(k) {
					badRects[rect] = true
					delete(badRects, k)
					inBounds = false
					break
				}
			}
			if !inBounds {
				continue
			}
			for y := minY; y <= maxY; y++ {
				if !inBounds {
					break
				}
				for x := minX; x <= maxX; x++ {
					if gridState[y][x] == Out {
						inBounds = false
						break
					}
				}
			}

			if !inBounds {
				badRects[rect] = true
			}

			if inBounds && area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func buildWall() {
	for i := 0; i < len(points); i++ {
		index := i
		a := points[index]
		if index+1 == len(points) {
			index = -1
		}
		b := points[index+1]
		if a.x == b.x {
			lo := a.y
			hi := b.y
			if a.y > b.y {
				lo, hi = hi, lo
			}
			for y := lo; y <= hi; y++ {
				gridState[y][a.x] = Wall
			}
		} else {
			lo := a.x
			hi := b.x
			if a.x > b.x {
				lo, hi = hi, lo
			}
			for x := lo; x <= hi; x++ {
				gridState[a.y][x] = Wall
			}
		}
	}
}

func floodFill() {
	q := make([]Point, 0, puzzleSize)
	q = append(q, Point{x: 0, y: 0})
	for len(q) > 0 {
		next := q[0]
		q = q[1:]
		if gridState[next.y][next.x] == Wall || gridState[next.y][next.x] == Out {
			continue
		}
		if next.x-1 >= 0 && gridState[next.y][next.x-1] == Empty {
			q = append(q, Point{next.x - 1, next.y})
			gridState[next.y][next.x-1] = Seen
		}
		if next.x+1 < len(gridState) && gridState[next.y][next.x+1] == Empty {
			q = append(q, Point{next.x + 1, next.y})
			gridState[next.y][next.x+1] = Seen
		}
		if next.y-1 >= 0 && gridState[next.y-1][next.x] == Empty {
			q = append(q, Point{next.x, next.y - 1})
			gridState[next.y-1][next.x] = Seen
		}
		if next.y+1 < len(gridState) && gridState[next.y+1][next.x] == Empty {
			q = append(q, Point{next.x, next.y + 1})
			gridState[next.y+1][next.x] = Seen
		}
		gridState[next.y][next.x] = Out
	}
}

func main() {
	for i := range gridState {
		gridState[i] = make([]State, puzzleSize)
	}
	readPoints()
	buildWall()
	floodFill()
	maxArea := getMaxRectangleArea()
	fmt.Println(maxArea)
}

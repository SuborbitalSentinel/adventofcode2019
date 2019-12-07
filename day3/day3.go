package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Run day 3 of advent
func Run(input <-chan string) {
	m := make(map[point][]int)

	wireNumber := 0
	for wirePath := range input {
		tokens := tokenizePath(wirePath)
		tracePath(tokens, wireNumber, m)
		wireNumber++
	}

	// point 0,0 will always have the wires connected
	delete(m, point{0, 0})

	minSeenDistance := math.MaxInt32
	for key, value := range m {
		if len(value) == 2 { // Assuming that there were be at most 2 values per point in space
			if value[0] != value[1] {
				distance := abs(key.X) + abs(key.Y)
				if distance < minSeenDistance {
					minSeenDistance = distance
				}
			}
		}
	}

	fmt.Println("Day3 -- Part1: ", minSeenDistance)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type point struct {
	X, Y int
}

func tokenizePath(input string) []string {
	return strings.Split(input, ",")
}

func parseToken(input string) (string, int) {
	direction := input[:1]
	distance, _ := strconv.Atoi(input[1:])
	return direction, distance
}

func goLeft(x *int, y int, distance int, wireNumber int, m map[point][]int) {
	for i := 1; i <= distance; i++ {
		*x--
		m[point{*x, y}] = append(m[point{*x, y}], wireNumber)
	}
}

func goRight(x *int, y int, distance int, wireNumber int, m map[point][]int) {
	for i := 1; i <= distance; i++ {
		*x++
		m[point{*x, y}] = append(m[point{*x, y}], wireNumber)
	}
}

func goUp(x int, y *int, distance int, wireNumber int, m map[point][]int) {
	for i := 1; i <= distance; i++ {
		*y++
		m[point{x, *y}] = append(m[point{x, *y}], wireNumber)
	}
}

func goDown(x int, y *int, distance int, wireNumber int, m map[point][]int) {
	for i := 1; i <= distance; i++ {
		*y--
		m[point{x, *y}] = append(m[point{x, *y}], wireNumber)
	}
}

func tracePath(tokens []string, wireNumber int, m map[point][]int) {
	m[point{0, 0}] = append(m[point{0, 0}], wireNumber)

	x, y := 0, 0
	for _, token := range tokens {
		direction, distance := parseToken(token)
		switch direction {
		case "R":
			goRight(&x, y, distance, wireNumber, m)
		case "L":
			goLeft(&x, y, distance, wireNumber, m)
		case "U":
			goUp(x, &y, distance, wireNumber, m)
		case "D":
			goDown(x, &y, distance, wireNumber, m)
		}
	}
}

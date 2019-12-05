package day1

import (
	"fmt"
	"strconv"
)

// Run day one of advent
func Run(input <-chan string) {
	simplifiedRequiredFuel := 0
	totalRequiredFuel := 0

	for line := range input {
		moduleMass, _ := strconv.Atoi(line)
		simplifiedRequiredFuel += CalculateSimplifiedRequiredFuelForModule(moduleMass)
		totalRequiredFuel += CalculateTotalRequiredFuelForModule(moduleMass, 0)
	}

	fmt.Println("Day1 -- Part1: ", simplifiedRequiredFuel)
	fmt.Println("Day1 -- Part2: ", totalRequiredFuel)
}

// CalculateSimplifiedRequiredFuelForModule problem 1 -- part 1
func CalculateSimplifiedRequiredFuelForModule(mass int) int {
	return (mass / 3) - 2
}

// CalculateTotalRequiredFuelForModule problem 1 -- part 2
func CalculateTotalRequiredFuelForModule(mass int, totalFuel int) int {
	if mass <= 0 {
		return totalFuel
	}
	amount := (mass / 3) - 2
	if amount >= 0 {
		totalFuel += amount
	}

	return CalculateTotalRequiredFuelForModule(amount, totalFuel)
}

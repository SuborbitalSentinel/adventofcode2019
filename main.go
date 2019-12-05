package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/millera023/adventofcode/day1"
)

func main() {
	Day1()
}

func Day1() {
	file, err := os.Open("day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	simplifiedRequiredFuel := 0
	totalRequiredFuel := 0
	for scanner.Scan() {
		moduleMass, _ := strconv.Atoi(scanner.Text())
		simplifiedRequiredFuel += day1.CalculateSimplifiedRequiredFuelForModule(moduleMass)
		totalRequiredFuel += day1.CalculateTotalRequiredFuelForModule(moduleMass, 0)
	}

	fmt.Println("Day1 -- Part1: ", simplifiedRequiredFuel)
	fmt.Println("Day1 -- Part2: ", totalRequiredFuel)
}

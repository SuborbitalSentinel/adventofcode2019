package day1

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

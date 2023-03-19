package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
// Думал что 37 надо строго находить 3 к десяткам и отдельно 7, но это оказалось неправильным.
func CalculateCost(carsCount int) uint {
	if carsCount%10 == 0 {
		return uint((carsCount / 10) * 95000)
	}
	return uint(carsCount/10*95000) + uint(carsCount%10*10000)
}

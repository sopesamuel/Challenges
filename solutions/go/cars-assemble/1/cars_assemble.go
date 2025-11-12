package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentRate := successRate / 100
    sumRate := percentRate * float64(productionRate)
    return float64(sumRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    perHourRate := CalculateWorkingCarsPerHour(productionRate, successRate)
    perMinuteRate := int(perHourRate) / 60
    return perMinuteRate
	
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	wholeCars := carsCount / 10 
    remCars := carsCount % 10

    realCost := wholeCars * 95000 + remCars * 10000
    return uint(realCost)
}

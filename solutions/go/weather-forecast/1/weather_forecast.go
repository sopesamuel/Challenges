// Package weather provides the program that solves the weather forecast issue.
package weather


var (
// CurrentCondition variable are used to store information.
	CurrentCondition string
// CurrentLocation variable are used to store information.
	CurrentLocation  string
)

// Forecast function does the entire forecast operation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

// Package weather provides tools to forecast the current weather condition of a city.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current city being forecasted.
var CurrentLocation string

// Forecast updates the current location and condition, then returns a formatted weather report.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

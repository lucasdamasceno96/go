// Package weather  provides tools about weather.
package weather

// CurrentCondition tells the current condition.
var CurrentCondition string

// CurrentLocation tells the current location.
var CurrentLocation string

// Forecast does his thing.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

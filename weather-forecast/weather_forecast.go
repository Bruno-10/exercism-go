// Package weather provides the forecast for a city.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast takes a city and a condition and returns the weather condition of
// the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

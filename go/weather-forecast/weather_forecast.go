// Package weather provides weather current conditions
// for a location.
package weather

// CurrentCondition represents current weather conditions.
var CurrentCondition string
// CurrentLocation represents the current location for forecasting.
var CurrentLocation string

// Forecast returns current weather conditions for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

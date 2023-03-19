// Package weather provides weather forecasting considering the city.
package weather

// CurrentCondition is the current weather to be forecasted.
var CurrentCondition string

// CurrentLocation is the current city concerned by the weather forecasting.
var CurrentLocation string

// Forecast provides weather forecasting considering the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

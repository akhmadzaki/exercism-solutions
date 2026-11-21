//Package weather provides the forecast weather for the current location.
package weather

//CurrentCondition represents the current condition of the location.
var CurrentCondition string
//CurrentLocation represents the current location.
var CurrentLocation string

//Forecast returns the current current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

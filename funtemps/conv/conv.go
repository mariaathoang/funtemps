package conv

import (
	"math"
)

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	return math.Round(value-32.0) * (5.0 / 9.0)
}

// Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 {
	return math.Round(value-32.0)*(5.0/9.0) + 273.15
}

// Konverterer Celsius til Farhenheit
func CelsiusToFarhenheit(value float64) float64 {
	return (value * (9.0 / 5.0)) + 32.0
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	return (value + 273.15)
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {
	return math.Round(value-273.15)*(9.0/5.0) + 32.0
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
	return math.Round(value - 273.15)
}

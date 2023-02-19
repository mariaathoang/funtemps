package main

import (
	"flag"
	"fmt"

	"conv_test.go/conv"
)

var fahr float64
var cels float64
var kelv float64
var out string

func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader farhenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

}

func main() {

	flag.Parse()

	// Konverterer Celsius
	var res float64
	if cels != 0 {
		if out == "F" {
			res = conv.CelsiusToFarhenheit(cels)
			fmt.Printf("%.1f°C er %.2f°F\n", cels, res)
		} else if out == "K" {
			res = conv.CelsiusToKelvin(cels)
			fmt.Printf("%.1f°C er %.2f°K\n", cels, res)
		}
	}
	// Konverterer Farhenheit
	var mov float64
	if fahr != 0 {
		if out == "C" {
			mov = conv.FarhenheitToCelsius(fahr)
			fmt.Printf("%.0f°F er %.0f°C\n", fahr, mov)
		} else if out == "K" {
			mov = conv.FarhenheitToKelvin(fahr)
			fmt.Printf("%.1f°F er %.2f°K\n", fahr, mov)
		}
	}

	// Konverterer Kelvin
	var ter float64
	if kelv != 0 {
		if out == "C" {
			ter = conv.KelvinToCelsius(kelv)
			fmt.Printf("%.2f°K er %.0f°C\n", kelv, ter)
		} else if out == "F" {
			ter = conv.KelvinToFarhenheit(kelv)
			fmt.Printf("%.2f°K er %.1f°F\n", kelv, ter)
		}
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

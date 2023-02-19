package conv

import (
	"testing"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134.01, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134.01, want: 329.82},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}
func TestCelciusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134.01},
	}

	for _, tc := range tests {
		got := CelsiusToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 329.82},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 134.01},
	}

	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 56.67},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

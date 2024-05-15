package stats

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty slice", []float64{}, 0},
		{"Single value", []float64{5}, 5},
		{"Positive values", []float64{1, 2, 3, 4, 5}, 3},
		{"Negative values", []float64{-1, -2, -3, -4, -5}, -3},
		{"Mixed values", []float64{-1, 2, -3, 4, -5}, -0.6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := Mean(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %f, but got %f", tc.expected, result)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty slice", []float64{}, 0},
		{"Single value", []float64{5}, 5},
		{"Even length, positive values", []float64{1, 2, 3, 4}, 2.5},
		{"Odd length, negative values", []float64{-1, -2, -3, -4, -5}, -3},
		{"Mixed values", []float64{-1, 2, -3, 4, -5, 6}, 0.5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := Median(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %f, but got %f", tc.expected, result)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty slice", []float64{}, 0},
		{"Single value", []float64{5}, 0},
		{"Positive values", []float64{1, 2, 3, 4, 5}, 2},
		{"Negative values", []float64{-1, -2, -3, -4, -5}, 2},
		{"Mixed values", []float64{-1, 2, -3, 4, -5}, 10.64},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := Variance(tc.input)
			if !floatEquals(result, tc.expected, 1e-9) {
				t.Errorf("Expected %f, but got %f", tc.expected, result)
			}
		})
	}
}

func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

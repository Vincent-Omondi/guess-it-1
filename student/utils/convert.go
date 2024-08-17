package utils

import (
	"math"
)

// Scaling factor for converting to integers
const scalingFactor = 1

// Function to convert float to int with scaling and overflow check
func FloatToInt(f float64) (int, bool) {
	scaledValue := f * scalingFactor
	if scaledValue > math.MaxInt64 || scaledValue < math.MinInt64 {
		return 0, true
	}
	return int(math.Round(scaledValue)), false
}

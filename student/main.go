package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"math-skills/stats"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var data []float64
	windowSize := 5 // Size of the moving window

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue // skip non-numeric values
		}
		data = append(data, value)

		// Ensure that we're only using the last 'windowSize' numbers
		if len(data) > windowSize {
			data = data[len(data)-windowSize:]
		}

		// Only calculate if we have enough data points
		if len(data) > 1 {
			mean := stats.Mean(data)
			stDev := math.Sqrt(stats.Variance(data))

			lowerBound := int(mean - 2*stDev)
			upperBound := int(mean + 2*stDev)

			fmt.Printf("%d %d\n", lowerBound, upperBound)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning input\n")
		return
	}
}

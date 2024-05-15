package stats

// Mean calculates the mean (average)
func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64 = 0

	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

// Median calculates the median
func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	data = BubbleSort(data)
	n := len(data)

	if n%2 == 0 {
		middle := n / 2
		return (data[middle-1] + data[middle]) / 2
	} else {
		middle := n / 2
		return data[middle]
	}
}

func BubbleSort(data []float64) []float64 {
	// sort the numbers using BubbleSort
	for i := 0; i < (len(data))-1; i++ {
		for j := 0; j < (len(data))-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

// Variance calculates the variance
func Variance(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	miu := Mean(data)
	var sqrs float64 = 0
	for _, value := range data {
		sqrs += (value - miu) * (value - miu)
	}
	return sqrs / float64(len(data))
}

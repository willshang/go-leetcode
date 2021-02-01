package main

import "math"

func main() {

}

// leetcode1093_大样本统计
func sampleStats(count []int) []float64 {
	maxValue, minValue := math.MinInt32, math.MaxInt32
	maxTime, maxTimeValue := 0, 0
	sum, total := 0, 0
	for i := 0; i < len(count); i++ {
		if count[i] > 0 {
			minValue = i
			break
		}
	}
	for i := len(count) - 1; i >= 0; i-- {
		if count[i] > 0 {
			maxValue = i
			break
		}
	}
	for i := 0; i < len(count); i++ {
		total = total + count[i]
		sum = sum + count[i]*i
		if count[i] > maxTime {
			maxTime = count[i]
			maxTimeValue = i
		}
	}
	a, b, temp := 0, 0, 0
	for i := 0; i < len(count); i++ {
		if temp <= (total-1)/2 && (total-1)/2 < temp+count[i] {
			a = i
		}
		if temp <= total/2 && total/2 < temp+count[i] {
			b = i
			break
		}
		temp = temp + count[i]
	}
	return []float64{
		float64(minValue),
		float64(maxValue),
		float64(sum) / float64(total),
		float64(a+b) / 2,
		float64(maxTimeValue),
	}
}

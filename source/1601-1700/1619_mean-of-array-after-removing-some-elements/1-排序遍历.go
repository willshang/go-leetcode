package main

import "sort"

func main() {

}

func trimMean(arr []int) float64 {
	temp := make([]float64, 0)
	for i := 0; i < len(arr); i++ {
		temp = append(temp, float64(arr[i]))
	}
	sort.Float64s(temp)
	var sum float64
	count := int(float64(len(arr)) * 0.05)
	for i := count; i < len(temp)-count; i++ {
		sum = sum + temp[i]
	}
	return sum / float64(len(arr)-2*count)
}

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findMinDifference([]string{"00:00", "23:59", "00:00"}))
}

// 剑指OfferII035.最小时间差
func findMinDifference(timePoints []string) int {
	arr := make([]int, 0)
	for i := 0; i < len(timePoints); i++ {
		value := getValue(timePoints[i])
		arr = append(arr, value)
	}
	sort.Ints(arr)
	res := math.MaxInt32
	arr = append(arr, arr[0]+1440)
	for i := 1; i < len(arr); i++ {
		if res > arr[i]-arr[i-1] {
			res = arr[i] - arr[i-1]
		}
	}
	return res
}

func getValue(str string) int {
	hour, _ := strconv.Atoi(str[:2])
	minute, _ := strconv.Atoi(str[3:])
	return hour*60 + minute
}

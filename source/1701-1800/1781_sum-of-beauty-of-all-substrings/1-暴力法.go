package main

import "math"

func main() {

}

// leetcode1781_所有子字符串美丽值之和
func beautySum(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		arr := [26]int{}
		arr[s[i]-'a']++
		for j := i + 2; j <= len(s); j++ {
			arr[s[j-1]-'a']++
			res = res + getCount(arr)
		}
	}
	return res
}

func getCount(arr [26]int) int {
	maxValue, minValue := math.MinInt32, math.MaxInt32
	for i := 0; i < 26; i++ {
		if arr[i] > 0 {
			if arr[i] > maxValue {
				maxValue = arr[i]
			}
			if arr[i] < minValue {
				minValue = arr[i]
			}
		}
	}
	return maxValue - minValue
}

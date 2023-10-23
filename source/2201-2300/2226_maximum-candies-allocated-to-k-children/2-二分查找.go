package main

func main() {

}

func maximumCandies(candies []int, k int64) int {
	sum := int64(0)
	maxValue := 0
	for i := 0; i < len(candies); i++ {
		sum = sum + int64(candies[i])
		maxValue = max(maxValue, candies[i])
	}
	if sum < k {
		return 0
	}
	left := 1
	right := maxValue + 1
	for left < right {
		mid := left + (right-left)/2
		total := check(candies, mid)
		if total >= k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func check(arr []int, target int) int64 {
	res := int64(0)
	for i := 0; i < len(arr); i++ {
		res = res + int64(arr[i]/target)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

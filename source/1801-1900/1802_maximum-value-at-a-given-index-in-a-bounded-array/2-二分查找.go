package main

func main() {

}

// leetcode1802_有界数组中指定下标处的最大值
func maxValue(n int, index int, maxSum int) int {
	res := 1
	leftTotal, rightTotal := index, n-index-1
	left, right := 1, maxSum+1
	for left < right {
		mid := left + (right-left)/2
		l := getTotal(mid, leftTotal)
		r := getTotal(mid, rightTotal)
		if l+r+mid <= maxSum {
			left = mid + 1
			res = mid
		} else {
			right = mid
		}
	}
	return res
}

func getTotal(high int, total int) int {
	need := high - 1
	if need >= total {
		return total * (need + high - total) / 2
	}
	return need*(1+need)/2 + total - need
}

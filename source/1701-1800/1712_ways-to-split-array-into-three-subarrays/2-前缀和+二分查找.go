package main

func main() {

}

func waysToSplit(nums []int) int {
	res := 0
	n := len(nums)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + nums[i]
	}
	for i := 1; i <= n-1; i++ {
		first := arr[i] // 第1个数
		if first*3 > arr[n] {
			break
		}
		left, right := i+1, n-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[n]-arr[mid] < arr[mid]-first {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		b := left
		left, right = i+1, n-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid]-first < first {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		a := left
		res = (res + b - a) % 1000000007
	}
	return res % 1000000007
}

package main

func main() {

}

func sumSubarrayMins(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	index := findMinValue(arr)
	left := index
	right := len(arr) - 1 - index
	// 1是自己，left是左边left个数+自己，right是右边right个数+自己，left*right是左边+自己+右边的个数
	count := 1 + left + right + left*right
	res := arr[index] * count % 1000000007
	res = res + sumSubarrayMins(arr[:index]) // 左边
	if index < len(arr)-1 {
		res = res + sumSubarrayMins(arr[index+1:]) // 右边
	}
	return res % 1000000007
}

func findMinValue(arr []int) int {
	minValue := arr[0]
	minIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}
	return minIndex
}

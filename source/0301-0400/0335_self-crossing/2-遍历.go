package main

func main() {

}

func isSelfCrossing(distance []int) bool {
	arr := append([]int{0, 0, 0, 0}, distance...)
	n := len(arr)
	i := 4
	// 圈变大
	for i < n && arr[i] > arr[i-2] {
		i++
	}
	if i == n {
		return false
	}
	if arr[i] >= arr[i-2]-arr[i-4] {
		arr[i-1] = arr[i-1] - arr[i-3]
	}
	i++
	// 圈变小
	for i < n && arr[i] < arr[i-2] {
		i++
	}
	if i == n {
		return false
	}
	return true
}

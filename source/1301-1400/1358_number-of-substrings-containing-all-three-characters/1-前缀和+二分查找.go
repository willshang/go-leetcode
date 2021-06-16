package main

func main() {

}

func numberOfSubstrings(s string) int {
	res := 0
	n := len(s)
	arr := make([][3]int, n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			arr[i+1][j] = arr[i][j]
		}
		value := int(s[i] - 'a')
		arr[i+1][value]++
	}
	for i := 0; i < n; i++ {
		left := i + 1
		right := n
		index := -1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid][0] > arr[i][0] &&
				arr[mid][1] > arr[i][1] &&
				arr[mid][2] > arr[i][2] {
				right = mid - 1
				index = mid
			} else {
				left = mid + 1
			}
		}
		if index != -1 {
			res = res + n - index + 1
		}
	}
	return res
}

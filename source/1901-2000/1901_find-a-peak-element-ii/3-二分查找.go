package main

func main() {

}

func findPeakGrid(mat [][]int) []int {
	n := len(mat)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		midValue, index := getMax(mat[mid])
		a, b := -1, -1
		if mid >= 1 {
			a, _ = getMax(mat[mid-1])
		}
		if mid <= n-2 {
			b, _ = getMax(mat[mid+1])
		}
		res, _ := getMax([]int{a, b, midValue})
		if res == midValue {
			return []int{mid, index}
		} else if res == a {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nil
}

func getMax(arr []int) (int, int) {
	res := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > res {
			index = i
			res = arr[i]
		}
	}
	return res, index
}

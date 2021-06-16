package main

func main() {

}

// leetcode1893_检查是否区域内所有整数都被覆盖
func isCovered(ranges [][]int, left int, right int) bool {
	arr := make([]bool, 51)
	for i := 0; i < len(ranges); i++ {
		a, b := ranges[i][0], ranges[i][1]
		for j := a; j <= b; j++ {
			arr[j] = true
		}
	}
	for i := left; i <= right; i++ {
		if arr[i] == false {
			return false
		}
	}
	return true
}

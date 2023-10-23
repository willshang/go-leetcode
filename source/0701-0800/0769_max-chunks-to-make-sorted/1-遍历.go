package main

func main() {

}

// leetcode769_最多能完成排序的块
func maxChunksToSorted(arr []int) int {
	res := 0
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if maxValue == i {
			res++
		}
	}
	return res
}

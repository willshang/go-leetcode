package main

func main() {

}

// leetcode1769_移动所有球到每个盒子所需的最小操作数
func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)
	right, rightCount := 0, 0
	for i := 0; i < n; i++ {
		if boxes[i] == '1' {
			right = right + i
			rightCount++
		}
	}
	left, leftCount := 0, 0
	for i := 0; i < n; i++ {
		res[i] = left + right
		if boxes[i] == '1' {
			leftCount++
			rightCount--
		}
		left = left + leftCount
		right = right - rightCount
	}
	return res
}

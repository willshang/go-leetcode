package main

func main() {

}

// leetcode1186_删除一次得到子数组最大和
func maximumSum(arr []int) int {
	n := len(arr)
	a := arr[0] // 不删
	b := 0      // 删除
	res := a    // 长度为1，不删除
	for i := 1; i < n; i++ {
		a, b = max(a+arr[i], arr[i]), max(b+arr[i], a)
		res = max(res, max(a, b))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

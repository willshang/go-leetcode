package main

func main() {

}

// leetcode1674_使数组互补的最少操作次数
func minMoves(nums []int, limit int) int {
	n := len(nums)
	arr := make([]int, 2*limit+2)
	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		// 1、首先[2,2*limit]都加2=>操作2次
		arr[2] = arr[2] + 2
		arr[2*limit+1] = arr[2*limit+1] - 2
		// 2、将[1+min(a,b),limit+max(a,b)] 减1=>操作1次
		arr[1+min(a, b)] = arr[1+min(a, b)] - 1
		arr[limit+max(a, b)+1] = arr[limit+max(a, b)+1] + 1
		// 3、将[a+b]减1，目标值=>操作0次
		arr[a+b] = arr[a+b] - 1
		arr[a+b+1] = arr[a+b+1] + 1
	}
	res := n
	sum := 0
	for i := 2; i <= 2*limit; i++ {
		sum = sum + arr[i]
		if res > sum {
			res = sum
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

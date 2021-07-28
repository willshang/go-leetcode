package main

func main() {

}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	n := len(students)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = calculate(students[i], mentors[j])
		}
	}
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = i
	}
	res := 0
	for {
		if temp == nil {
			break
		}
		sum := 0
		for i := 0; i < n; i++ {
			sum = sum + arr[i][temp[i]]
		}
		if sum > res {
			res = sum
		}
		temp = nextPermutation(temp)
	}
	return res
}

// leetcode31.下一个排列
func nextPermutation(nums []int) []int {
	n := len(nums)
	left := n - 2
	// 以12385764为例，从后往前找到5<7 的升序情况，目标值为左边的数5
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left == -1 { // 排完了，下一个返回nil结束
		return nil
	}
	right := n - 1
	// 从后往前，找到第一个大于目标值的数，如6>5，然后交换
	for right >= 0 && nums[right] <= nums[left] {
		right--
	}
	nums[left], nums[right] = nums[right], nums[left]
	count := 0
	// 后面是降序状态，让它变为升序
	for i := left + 1; i <= (left+1+n-1)/2; i++ {
		nums[i], nums[n-1-count] = nums[n-1-count], nums[i]
		count++
	}
	return nums
}

func calculate(a, b []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res++
		}
	}
	return res
}

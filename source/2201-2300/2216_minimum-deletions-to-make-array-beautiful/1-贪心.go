package main

func main() {

}

func minDeletion(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; {
		j := i + 1
		for j < n && nums[i] == nums[j] { // 找到后面1个不相同的组成1对，然后开始找下一对
			j++
		}
		if j < n {
			res = res + 2
		}
		i = j + 1
	}
	return n - res
}

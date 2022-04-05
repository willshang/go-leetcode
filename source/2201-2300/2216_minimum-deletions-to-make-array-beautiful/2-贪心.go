package main

func main() {

}

// leetcode2216_美化数组的最少删除数
func minDeletion(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			res++
		} else {
			i++
		}
	}
	return res + (n-res)%2 // 剩下奇数个要+1
}

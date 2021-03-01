package main

func main() {

}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	return target == 1
	// return dfs(nums, k, target,make([]bool, len(nums)), 0, 0)
}
func dfs(nums []int, k int) {

}

package main

func main() {

}

// leetcode1031_两个非重叠子数组的最大和
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	res := 0
	n := len(nums)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + nums[i]
	}
	for i := firstLen; i <= n; i++ { // 枚举L的起始点
		L := arr[i] - arr[i-firstLen]
		M := 0
		for j := secondLen; j <= i-firstLen; j++ { // M在L左边的情况
			M = max(M, arr[j]-arr[j-secondLen])
		}
		for j := n; j >= i+secondLen; j-- { // M在L右边的情况
			M = max(M, arr[j]-arr[j-secondLen])
		}
		res = max(res, L+M)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

func main() {

}

// leetcode2121_相同元素的间隔之和
func getDistances(arr []int) []int64 {
	n := len(arr)
	res := make([]int64, n)
	m := make(map[int][]int) // 对相同数字进行分组
	for i := 0; i < n; i++ {
		m[arr[i]] = append(m[arr[i]], i)
	}
	for _, v := range m { // 分组
		arr := getSumAbsoluteDifferences(v)
		for i := 0; i < len(v); i++ {
			res[v[i]] = int64(arr[i])
		}
	}
	return res
}

// leetcode1685.有序数组中差绝对值之和
func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	arr := make([]int, 0)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
		arr = append(arr, sum)
	}
	res = append(res, sum-n*nums[0])
	for i := 1; i < n; i++ {
		left := nums[i]*i - arr[i-1]              // 左边和
		right := (sum - arr[i]) - nums[i]*(n-1-i) // 右边和
		res = append(res, left+right)
	}
	return res
}

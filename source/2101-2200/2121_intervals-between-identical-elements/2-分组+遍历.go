package main

func main() {

}

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
	right := 0 // 右边和
	left := 0  // 左边和
	for i := 1; i < n; i++ {
		right = right + (nums[i] - nums[0])
	}
	res = append(res, right)
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		left = left + diff*i
		right = right - diff*(n-i)
		res = append(res, left+right)
	}
	return res
}

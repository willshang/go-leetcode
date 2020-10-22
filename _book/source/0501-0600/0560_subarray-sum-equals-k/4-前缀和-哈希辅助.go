package main

import "fmt"

func main() {
	//fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	//fmt.Println(subarraySum([]int{2, 5, 2, 5}, 7)) // 3
	fmt.Println(subarraySum([]int{0, 1, -1, 1, -1}, 0)) // 3
}

func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int][]int)
	m[0] = []int{-1} // 保证第一个k的存在
	sum := 0
	// sum[i:j]= sum[0:j]-sum[0:i]，把sum[i:j]设为k，
	// 于是可以转化为sum[0:j]-k=sum[0:i]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := m[sum-k]; ok {
			res = res + len(m[sum-k])
			// 输出满足条件的子数组下标
			// for _, v := range m[sum-k] {
			//	fmt.Println(v+1, i)
			// }
		}
		m[sum] = append(m[sum], i)
	}
	return res
}

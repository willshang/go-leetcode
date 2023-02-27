package main

func main() {

}

func minimumOperations(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			m[nums[i]] = true // 统计不为0的数字去重个数
		}
	}
	return len(m)
}

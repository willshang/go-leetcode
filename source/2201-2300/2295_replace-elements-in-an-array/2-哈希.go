package main

func main() {

}

func arrayChange(nums []int, operations [][]int) []int {
	m := make(map[int]int)
	for i := len(operations) - 1; i >= 0; i-- {
		a, b := operations[i][0], operations[i][1]
		if v, ok := m[b]; ok {
			b = v // 在后面出现过，使用最后的结果
		}
		m[a] = b
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			nums[i] = v
		}
	}
	return nums
}

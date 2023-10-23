package main

func main() {

}

// leetcode2295_替换数组中的元素
func arrayChange(nums []int, operations [][]int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(operations); i++ {
		a, b := operations[i][0], operations[i][1]
		index := m[a]
		delete(m, a)
		m[b] = index
		nums[index] = b
	}
	return nums
}

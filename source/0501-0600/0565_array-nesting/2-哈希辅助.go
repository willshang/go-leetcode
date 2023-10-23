package main

func main() {

}

// leetcode565_数组嵌套
func arrayNesting(nums []int) int {
	m := make(map[int]bool)
	res := 0
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == true {
			continue
		}
		count := 0
		cur := i
		for {
			count++
			m[cur] = true
			cur = nums[cur]
			if m[cur] == true {
				break
			}
		}
		if count > res {
			res = count
		}
	}
	return res
}

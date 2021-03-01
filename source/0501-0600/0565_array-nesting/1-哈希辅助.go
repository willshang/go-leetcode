package main

func main() {

}

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
			if cur == i {
				break
			}
		}
		if count > res {
			res = count
		}
	}
	return res
}

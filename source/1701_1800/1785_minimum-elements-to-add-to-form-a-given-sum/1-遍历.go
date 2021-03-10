package main

func main() {

}

// leetcode1785_构成特定和需要添加的最少元素
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	left := abs(goal - sum)
	a := left / limit
	b := left % limit
	if b != 0 {
		a = a + 1
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

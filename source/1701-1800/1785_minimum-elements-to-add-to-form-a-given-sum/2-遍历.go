package main

func main() {

}

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	res := abs(goal - sum)
	return (res - 1 + limit) / limit
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

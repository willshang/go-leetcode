package main

func main() {

}

// leetcode2154_将找到的值乘以2
func findFinalValue(nums []int, original int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for m[original] == true {
		original = original * 2
	}
	return original
}

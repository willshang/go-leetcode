package main

func main() {

}

// leetcode2170_使数组变成交替数组的最少操作数
func minimumOperations(nums []int) int {
	m := [2]map[int]int{}
	for i := 0; i < 2; i++ {
		m[i] = make(map[int]int)
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			m[0][nums[i]]++
		} else {
			m[1][nums[i]]++
		}
	}
	a := getMaxValue(m[0])
	b := getMaxValue(m[1])
	if a[0] == b[0] {
		return len(nums) - max(a[1]+b[3], a[3]+b[1])
	}
	return len(nums) - a[1] - b[1]
}

func getMaxValue(m map[int]int) []int {
	firstValue, firstIndex := 0, 0
	secondValue, secondIndex := 0, 0
	for k, v := range m {
		if v > firstValue {
			secondIndex, secondValue = firstIndex, firstValue
			firstIndex, firstValue = k, v
		} else if v > secondValue {
			secondIndex, secondValue = k, v
		}
	}
	return []int{firstIndex, firstValue, secondIndex, secondValue}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

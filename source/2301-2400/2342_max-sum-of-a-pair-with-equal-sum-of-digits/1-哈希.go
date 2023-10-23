package main

func main() {

}

func maximumSum(nums []int) int {
	res := -1
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v := calculate(nums[i])
		if m[v] > 0 { // 存在值
			res = max(res, m[v]+nums[i]) // 更新最大值
		}
		m[v] = max(m[v], nums[i])
	}
	return res
}

func calculate(a int) int {
	res := 0
	for a > 0 {
		res = res + a%10
		a = a / 10
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

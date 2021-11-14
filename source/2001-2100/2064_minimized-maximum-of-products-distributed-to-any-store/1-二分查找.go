package main

func main() {

}

// leetcode2064_分配给商店的最多商品的最小值
func minimizedMaximum(n int, quantities []int) int {
	left, right := 1, 100000
	for left < right {
		mid := left + (right-left)/2
		if judge(quantities, mid) <= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(quantities []int, per int) int {
	res := 0
	for i := 0; i < len(quantities); i++ {
		res = res + quantities[i]/per
		if quantities[i]%per != 0 {
			res++
		}
	}
	return res
}

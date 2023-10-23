package main

func main() {

}

func maximumGroups(grades []int) int {
	n := len(grades)
	res := 1 // 至少可以分为1组
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2 // 分为mid组:1+2+3+...+mid（多余的分在最后一组）
		total := (1 + mid) * mid / 2 // 分为mid组至少需要的人数
		if total > n {
			right = mid - 1
		} else {
			res = mid
			left = mid + 1
		}
	}
	return res
}

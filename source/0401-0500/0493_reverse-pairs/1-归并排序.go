package main

func main() {

}

// leetcode493_翻转对
func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	a := append([]int{}, nums[:n/2]...)
	b := append([]int{}, nums[n/2:]...)
	res := reversePairs(a) + reversePairs(b) // 递归后，a、b分别有序
	i, j := 0, 0
	for i = 0; i < len(a); i++ {
		for j < len(b) && a[i] > 2*b[j] { // 统计
			j++
		}
		res = res + j
	}
	i, j = 0, 0
	for k := 0; k < len(nums); k++ { // 2个有序数组合并
		if i < len(a) && (j == len(b) || a[i] <= b[j]) {
			nums[k] = a[i]
			i++
		} else {
			nums[k] = b[j]
			j++
		}
	}
	return res
}

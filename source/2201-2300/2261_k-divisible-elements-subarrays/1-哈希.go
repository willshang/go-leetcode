package main

func main() {

}

func countDistinct(nums []int, k int, p int) int {
	m := make(map[[200]int]bool)
	n := len(nums)
	for i := 0; i < n; i++ { // 左边界
		count := 0
		for j := i + 1; j <= n; j++ { // 右边界
			if nums[j-1]%p == 0 {
				count++
			}
			if count <= k {
				temp := [200]int{}
				for k := i; k < j; k++ {
					temp[k-i] = nums[k]
				}
				m[temp] = true
			}
		}
	}
	return len(m)
}

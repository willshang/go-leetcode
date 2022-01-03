package main

import "sort"

func main() {

}

func recoverArray(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[0] || (nums[i]-nums[0])%2 == 1 || nums[i] == nums[i-1] {
			continue
		}
		visited := make([]bool, n)
		visited[0], visited[i] = true, true
		res := make([]int, 0)
		k := (nums[i] - nums[0]) / 2
		res = append(res, nums[0]+k)
		left := 1
		right := i
		for j := 2; j <= n/2; j++ { // 还需要遍历n/2-1次
			for visited[left] == true {
				left++
			}
			for right < n && (visited[right] == true || nums[right]-nums[left] != 2*k) {
				right++
			}
			if right == n {
				break
			}
			res = append(res, nums[left]+k)
			visited[left], visited[right] = true, true
		}
		if len(res) == n/2 {
			return res
		}
	}
	return nil
}

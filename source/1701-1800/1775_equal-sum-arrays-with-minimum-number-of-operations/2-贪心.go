package main

func main() {

}

// leetcode1775_通过最少操作次数使数组的和相等
func minOperations(nums1 []int, nums2 []int) int {
	aMin, bMin := len(nums1), len(nums2)
	aMax, bMax := aMin*6, bMin*6
	if aMin > bMax || aMax < bMin {
		return -1
	}
	a, b := 0, 0
	arrA, arrB := [7]int{}, [7]int{}
	for i := 0; i < len(nums1); i++ {
		a = a + nums1[i]
		arrA[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		b = b + nums2[i]
		arrB[nums2[i]]++
	}
	if a == b {
		return 0
	}
	if b > a {
		arrA, arrB = arrB, arrA
		a, b = b, a
	}
	arr := make([]int, 6) // 存储
	for i := 1; i < 6; i++ {
		arr[i-1] = arr[i-1] + arrA[7-i] + arrB[i]
	}
	res, total := 0, a-b
	for i := 0; i < len(arr); i++ {
		diff := 5 - i
		if total-diff*arr[i] > 0 {
			total = total - diff*arr[i]
			res = res + arr[i]
		} else {
			if total%diff == 0 {
				res = res + total/diff
			} else {
				res = res + total/diff + 1
			}
			return res
		}
	}
	return res
}

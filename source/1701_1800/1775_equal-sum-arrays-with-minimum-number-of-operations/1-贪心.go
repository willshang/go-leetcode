package main

func main() {

}

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
	arr := make([]int, 0) // 存储
	for i := 1; i < 6; i++ {
		if arrA[7-i] > arrB[i] {
			arr = append(arr, arrA[7-i], arrB[i])
		} else {
			arr = append(arr, arrB[i], arrA[7-i])
		}
	}
	res := 0
	total := a - b
	for i := 0; i < len(arr); i++ {
		diff := 6 - (i+2)/2
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

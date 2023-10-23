package main

func main() {

}

// leetcode1865_找出和为指定值的下标对
type FindSumPairs struct {
	nums1, nums2 []int
	m            map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]]++
	}
	return FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		m:     m,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.m[this.nums2[index]]--
	this.nums2[index] = this.nums2[index] + val
	this.m[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	res := 0
	for i := 0; i < len(this.nums1); i++ {
		res = res + this.m[tot-this.nums1[i]]
	}
	return res
}

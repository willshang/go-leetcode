package main

import "fmt"

func main() {
	a := []int{1,2,3,0,0,0}
	merge(a,3,[]int{},0)
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	temp := make([]int,m)
	copy(temp,nums1)

	if n == 0 {
		return
	}
	first,second := 0,0
	for i := 0; i < len(nums1); i++{
		if second >= n{
			nums1[i] = temp[first]
			first++
			continue
		}
		if first >= m{
			nums1[i] = nums2[second]
			second++
			continue
		}
		if temp[first] < nums2[second]{
			nums1[i] = temp[first]
			first++
		}else {
			nums1[i] = nums2[second]
			second++
		}
	}
}
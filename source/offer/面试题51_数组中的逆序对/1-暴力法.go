package main

import "fmt"

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}

func reversePairs(nums []int) int {
	res := 0
	if len(nums) <= 1 {
		return res
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	res = InversePairsCore(&nums, &temp, 0, len(nums)-1)
	return res
}

func InversePairsCore(num, temp *[]int, start, end int) int {
	fmt.Println(num, temp)
	if start == end {
		temp[start] = num[start]
		return 0
	}
	mid := (end - start) / 2
	left := InversePairsCore(num, temp, start, start+mid)
	right := InversePairsCore(num, temp, start+mid+1, end)
	i := start + mid
	j := end
	indexCopy := end
	count := 0
	for i >= start && j >= start+mid+1 {
		if num[i] > num[j] {
			temp[indexCopy] = num[i]
			indexCopy--
			i--
			count = count + j - start - mid
		} else {
			temp[indexCopy] = num[i]
			indexCopy--
			i--
		}
	}
	i--
	for ; i >= start; i-- {
		temp[indexCopy] = num[i]
		indexCopy--
	}
	j--
	for ; j >= start+mid+1; j-- {
		temp[indexCopy] = num[j]
		indexCopy--
	}
	return left + right + count
}

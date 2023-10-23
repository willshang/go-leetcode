package main

import "sort"

func main() {

}

// leetcode1054_距离相等的条形码
type Node struct {
	value int
	num   int
}

func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(barcodes); i++ {
		m[barcodes[i]]++
	}
	arr := make([]Node, 0)
	for k, v := range m {
		arr = append(arr, Node{
			k,
			v,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].num > arr[j].num
	})
	res := make([]int, len(barcodes))
	index := 0
	// 先偶后奇
	for i := 0; i < 2; i++ {
		for j := i; j < len(barcodes); j = j + 2 {
			if arr[index].num == 0 {
				index++
			}
			res[j] = arr[index].value
			arr[index].num--
		}
	}
	return res
}

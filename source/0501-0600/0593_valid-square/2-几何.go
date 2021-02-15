package main

import "sort"

func main() {

}

// leetcode593_有效的正方形
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	arr := make([]int, 0)
	arr = append(arr, getDis(p1, p2))
	arr = append(arr, getDis(p1, p3))
	arr = append(arr, getDis(p1, p4))
	arr = append(arr, getDis(p2, p3))
	arr = append(arr, getDis(p2, p4))
	arr = append(arr, getDis(p3, p4))
	sort.Ints(arr)
	return arr[0] > 0 && arr[0] == arr[3] && arr[4] == arr[5] && arr[0]*2 == arr[4]
}

func getDis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

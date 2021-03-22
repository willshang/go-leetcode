package main

import "fmt"

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
	//fmt.Println(countTriplets([]int{1, 2, 3, 4, 2}))
}

func countTriplets(arr []int) int {
	res := 0
	sumM := make(map[int]int)
	countM := make(map[int]int)
	countM[0] = 1
	temp := 0
	for i := 0; i < len(arr); i++ {
		temp = temp ^ arr[i]
		if countM[temp] > 0 {
			// a[i]^...a[j-1]^a[j]^...a[k] = 0，则j可以取i+1、i+2、...、..k
			// 相同异或结果，分别出现在下标[a,b,c,d]
			// 则[a,d]有d-a-1个满足条件的
			// sum = (d-a-1)+(d-b-1)+(d-c-1)
			// ==> nd - [(a+1) + (b+1) + (c+1)]
			// 同理得[a,b], [a,c]
			res = res + i*countM[temp] - sumM[temp]
		}
		countM[temp]++
		sumM[temp] = sumM[temp] + (i + 1)
	}
	return res
}

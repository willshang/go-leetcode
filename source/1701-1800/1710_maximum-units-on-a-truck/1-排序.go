package main

import "sort"

func main() {

}

// leetcode1710_卡车上的最大单元数
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	for i := 0; i < len(boxTypes); i++ {
		if boxTypes[i][0] <= truckSize {
			res = res + boxTypes[i][0]*boxTypes[i][1]
		} else {
			res = res + truckSize*boxTypes[i][1]
			break
		}
		truckSize = truckSize - boxTypes[i][0]
	}
	return res
}

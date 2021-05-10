package main

func main() {

}

// leetcode1854_人口最多的年份
func maximumPopulation(logs [][]int) int {
	arr := [101]int{}
	for i := 0; i < len(logs); i++ {
		a, b := logs[i][0], logs[i][1]
		for j := a; j < b; j++ {
			arr[j-1950]++
		}
	}
	res := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > count {
			count = arr[i]
			res = i + 1950
		}
	}
	return res
}

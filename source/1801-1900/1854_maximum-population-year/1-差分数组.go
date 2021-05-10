package main

func main() {

}

func maximumPopulation(logs [][]int) int {
	arr := [101]int{}
	for i := 0; i < len(logs); i++ {
		a, b := logs[i][0], logs[i][1]
		arr[a-1950]++
		arr[b-1950]--
	}
	res := 0
	sum := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum > count {
			count = sum
			res = i + 1950
		}
	}
	return res
}

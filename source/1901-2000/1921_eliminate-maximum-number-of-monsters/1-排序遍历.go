package main

import "sort"

func main() {

}

// leetcode1921_消灭怪物的最大数量
func eliminateMaximum(dist []int, speed []int) int {
	res := 0
	temp := make([]float64, len(dist))
	for i := 0; i < len(dist); i++ {
		temp[i] = float64(dist[i]) / float64(speed[i])
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	for i := 0; i < len(dist); i++ {
		if float64(i) < temp[i] {
			res++
		} else {
			break
		}
	}
	return res
}

package main

import "fmt"

func main() {
	//fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{20, 40}))
}

func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	for i := 0; i < len(time); i++ {
		m[time[i]%60]++
	}
	res := 0
	for key, value := range m {
		if key == 0 || key == 30 {
			res = res + (value-1)*value/2
		} else {
			if v, ok := m[60-key]; ok && v > 0 {
				res = res + v*value
				m[key] = 0
				m[60-key] = 0
			}
		}
	}
	return res
}

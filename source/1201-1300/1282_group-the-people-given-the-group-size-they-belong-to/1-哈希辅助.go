package main

func main() {

}

// leetcode1282_用户分组
func groupThePeople(groupSizes []int) [][]int {
	res := make([][]int, 0)
	m := make(map[int][]int)
	for i := 0; i < len(groupSizes); i++ {
		m[groupSizes[i]] = append(m[groupSizes[i]], i)
	}
	for k, v := range m {
		for i := 0; i < len(v); i = i + k {
			res = append(res, v[i:i+k])
		}
	}
	return res
}

package main

func main() {

}

func groupThePeople(groupSizes []int) [][]int {
	res := make([][]int, 0)
	m := make(map[int][]int)
	for i := 0; i < len(groupSizes); i++ {
		m[groupSizes[i]] = append(m[groupSizes[i]], i)
		if groupSizes[i] == len(m[groupSizes[i]]) {
			res = append(res, m[groupSizes[i]])
			m[groupSizes[i]] = []int{}
		}
	}
	return res
}

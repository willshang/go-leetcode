package main

func main() {

}

// leetcode1817_查找用户活跃分钟数
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	m := make(map[int]map[int]int)
	for i := 0; i < len(logs); i++ {
		a, b := logs[i][0], logs[i][1]
		if _, ok := m[a]; ok == false {
			m[a] = make(map[int]int)
		}
		m[a][b]++
	}
	res := make([]int, k)
	for _, v := range m {
		value := len(v)
		res[value-1]++
	}
	return res
}

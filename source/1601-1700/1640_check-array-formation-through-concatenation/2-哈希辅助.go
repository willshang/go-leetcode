package main

func main() {

}

// leetcode1640_能否连接形成数组
func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i
	}
	for i := 0; i < len(pieces); i++ {
		if _, ok := m[pieces[i][0]]; !ok {
			return false
		}
		for k := 0; k < len(pieces[i])-1; k++ {
			if m[pieces[i][k+1]]-m[pieces[i][k]] != 1 {
				return false
			}
		}
	}
	return true
}

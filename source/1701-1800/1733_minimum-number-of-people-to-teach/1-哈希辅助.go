package main

func main() {

}

// leetcode1733_需要教语言的最少人数
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := make(map[int]map[int]int)
	for i := 0; i < len(languages); i++ {
		a := i + 1
		m[a] = make(map[int]int)
		for j := 0; j < len(languages[i]); j++ {
			b := languages[i][j]
			m[a][b] = 1
		}
	}
	need := make(map[int]bool) // 需要沟通的人列表
	for i := 0; i < len(friendships); i++ {
		a, b := friendships[i][0], friendships[i][1]
		flag := false
		for j := 1; j <= n; j++ {
			if m[a][j] == 1 && m[b][j] == 1 {
				flag = true
				break
			}
		}
		if flag == false {
			need[a] = true
			need[b] = true
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		count := 0
		for k := range need {
			if m[k][i] == 1 {
				count++
			}
		}
		if count > res {
			res = count
		}
	}
	return len(need) - res
}

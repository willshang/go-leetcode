package main

func main() {

}

// leetcode2244_完成所有任务需要的最少轮数
func minimumRounds(tasks []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
	}
	for _, v := range m {
		if v == 1 { // 1个直接返回
			return -1
		}
		if v%3 == 0 {
			res = res + v/3
		} else {
			res = res + v/3 + 1 // %v=1 拆成2+2；%v=2，拆成3+2
		}
	}
	return res
}

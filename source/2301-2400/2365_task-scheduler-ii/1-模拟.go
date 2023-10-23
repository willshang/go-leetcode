package main

func main() {

}

func taskSchedulerII(tasks []int, space int) int64 {
	res := int64(0)
	m := make(map[int]int64)
	for i := 0; i < len(tasks); i++ {
		if m[tasks[i]] == 0 {
			res = res + 1 // 没出现+1
		} else {
			prev := m[tasks[i]]
			res = max(res+1, prev+int64(space)+1) // 出现：取较大
		}
		m[tasks[i]] = res
	}
	return res
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

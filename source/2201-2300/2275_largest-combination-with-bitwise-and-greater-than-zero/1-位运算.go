package main

func main() {

}

func largestCombination(candidates []int) int {
	arr := [32]int{}
	for i := 0; i < len(candidates); i++ {
		v := candidates[i]
		j := 0
		for v > 0 {
			if v%2 == 1 {
				arr[j]++
			}
			v = v / 2
			j++
		}
	}
	res := 0
	for i := 0; i < 32; i++ {
		res = max(res, arr[i]) // 计算二进制1出现最多的次数
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

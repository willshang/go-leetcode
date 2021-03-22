package main

func main() {

}

func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)
	pre := make([]int, n)
	count, sum := 0, 0
	for i := 0; i < n; i++ {
		pre[i] = sum
		if boxes[i] == '1' {
			count++
		}
		sum = sum + count
	}
	suf := make([]int, n)
	count, sum = 0, 0
	for i := n - 1; i >= 0; i-- {
		suf[i] = sum
		if boxes[i] == '1' {
			count++
		}
		sum = sum + count
	}
	for i := 0; i < n; i++ {
		res[i] = pre[i] + suf[i]
	}
	return res
}

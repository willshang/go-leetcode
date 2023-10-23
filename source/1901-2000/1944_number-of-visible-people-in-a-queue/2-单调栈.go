package main

func main() {

}

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	stack := make([]int, 0) // 递减栈
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] > heights[stack[len(stack)-1]] {
			res[i]++
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i]++ // 非空，还可以看到一个人
		}
		stack = append(stack, i)
	}
	return res
}

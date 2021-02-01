package main

func main() {

}

func longestWPI(hours []int) int {
	arr := make([]int, 0)
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			arr = append(arr, 1)
		} else {
			arr = append(arr, -1)
		}
	}
	temp := make([]int, len(hours)+1)
	for i := 1; i <= len(hours); i++ {
		temp[i] = temp[i-1] + arr[i-1]
	}
	stack := make([]int, 0)
	// 单调栈
	for i := 0; i < len(temp); i++ {
		if len(stack) == 0 || temp[stack[len(stack)-1]] > temp[i] {
			stack = append(stack, i)
		}
	}
	res := 0
	for i := len(temp) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			break
		}
		for len(stack) > 0 && temp[i] > temp[stack[len(stack)-1]] {
			if i-stack[len(stack)-1] > res {
				res = i - stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

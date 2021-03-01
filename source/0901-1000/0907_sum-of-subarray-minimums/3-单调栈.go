package main

func main() {

}

func sumSubarrayMins(arr []int) int {
	res := 0
	stack := make([][2]int, 0) // 递增栈
	sum := 0
	for i := 0; i < len(arr); i++ {
		count := 1
		for len(stack) > 0 && arr[i] < stack[len(stack)-1][0] { // 小于栈顶
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum = sum - prev[0]*prev[1]
			count = count + prev[1]
		}
		stack = append(stack, [2]int{arr[i], count})
		sum = sum + arr[i]*count
		res = (res + sum) % 1000000007
	}
	return res % 1000000007
}

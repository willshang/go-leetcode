package main

func main() {

}

func numSubarraysWithSum(A []int, S int) int {
	res := 0
	sum := 0
	left := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
		if sum > S { // 左指针右移
			for left < i && sum != S {
				sum = sum - A[left]
				left++
			}
		}
		if S == sum {
			tempSum := sum
			j := left
			for j <= i && tempSum == S { // 以该节点为结尾有多少个子数组sum==S
				res++
				tempSum = tempSum - A[j]
				j++
			}
		}
	}
	return res
}

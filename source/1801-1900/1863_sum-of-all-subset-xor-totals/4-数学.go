package main

func main() {

}

func subsetXORSum(nums []int) int {
	// 对于任意一位，2^n个子集异或结果中
	// 如果nums中任何一个数字在这一位上都是0，则任何一个子集的异或结果在这一位上都是0
	// 如果nums中有一个数字在这一位上是1，则所有子集异或结果中在这一位上，一半是0，一半是1
	n := len(nums)
	temp := 0
	for i := 0; i < n; i++ {
		temp = temp | nums[i]
	}
	return temp << (n - 1)
}

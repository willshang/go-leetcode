package main

func main() {

}

// leetcode1073_负二进制数相加
func addNegabinary(arr1 []int, arr2 []int) []int {
	res := make([]int, 1005)
	last := 1004
	i := len(arr1) - 1
	j := len(arr2) - 1
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry = carry + arr1[i]
			i--
		}
		if j >= 0 {
			carry = carry + arr2[j]
			j--
		}
		// 进位处理：
		// 进位可能：-1（0+0-1）、0、1、2、3（1+1+1）
		// 进位计算：-1 => 1; 0、1 => 0; 2、3=>-1
		res[last] = abs(carry) % 2
		if carry >= 0 {
			carry = -carry / 2
		} else {
			carry = 1
		}
		last--
	}
	for last < len(res)-1 && res[last] == 0 { // 消除多余的0，最后1个0不消除
		last++
	}
	return res[last:]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

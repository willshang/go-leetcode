package main

func main() {

}

// leetcode1653_使字符串平衡的最少删除次数
func minimumDeletions(s string) int {
	res := 0
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			stack = append(stack, 'b')
		} else {
			// 遇到'ba'则删掉，并次数+1
			if len(stack) > 0 {
				res++
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res
}

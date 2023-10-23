package main

func main() {

}

var mod = 1000000007

func uniqueLetterString(s string) int {
	res := 0
	n := len(s)
	arr := [26][]int{} // 记录每个字符的所有下标列表
	for i := 0; i < n; i++ {
		index := int(s[i] - 'A')
		arr[index] = append(arr[index], i)
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < len(arr[i]); j++ {
			var prev, next int
			cur := arr[i][j]
			if j == 0 {
				prev = -1
			} else {
				prev = arr[i][j-1]
			}
			if j == len(arr[i])-1 {
				next = n
			} else {
				next = arr[i][j+1]
			}
			res = (res + (cur-prev)*(next-cur)) % mod
		}
	}
	return res
}

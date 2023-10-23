package main

func main() {

}

var mod = 1000000007

func uniqueLetterString(s string) int {
	res := 0
	n := len(s)
	arr := [26][2]int{} // 记录每个字符的最后2个字符下标
	for i := 0; i < 26; i++ {
		arr[i] = [2]int{-1, -1}
	}
	sum := 0
	for i := 0; i < n; i++ {
		index := int(s[i] - 'A')
		cur := i - arr[index][1]
		prev := arr[index][1] - arr[index][0]
		sum = sum + cur - prev
		res = (res + sum) % mod
		// 更新下标
		arr[index][0] = arr[index][1]
		arr[index][1] = i
	}
	return res
}

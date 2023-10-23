package main

func main() {

}

var mod = 1000000007
var cur [26]int

func uniqueLetterString(s string) int {
	res := 0
	n := len(s)
	cur = [26]int{}    // 记录每个字符的下标进度
	arr := [26][]int{} // 记录每个字符的所有下标列表
	for i := 0; i < n; i++ {
		index := int(s[i] - 'A')
		arr[index] = append(arr[index], i)
	}
	sum := 0 // 26种字符，作为唯一字符的总次数
	for i := 0; i < 26; i++ {
		arr[i] = append(arr[i], n, n) // 补2个n
		sum = sum + getDiff(arr, i)
	}
	for i := 0; i < n; i++ {
		res = (res + sum) % mod
		index := int(s[i] - 'A')
		prev := getDiff(arr, index)
		cur[index]++
		sum = sum + getDiff(arr, index) - prev // 更新次数
	}
	return res
}

func getDiff(arr [26][]int, index int) int { // 第i+1个和第i个下标的距离
	i := cur[index]
	return arr[index][i+1] - arr[index][i]
}

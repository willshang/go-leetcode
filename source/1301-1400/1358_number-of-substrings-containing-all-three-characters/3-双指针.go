package main

func main() {

}

// leetcode1358_包含所有三种字符的子字符串数目
func numberOfSubstrings(s string) int {
	res := 0
	n := len(s)
	arr := [3]int{}
	var value int
	i := 0
	for j := 0; j < n; j++ {
		value = int(s[j] - 'a')
		arr[value]++
		for arr[0] > 0 && arr[1] > 0 && arr[2] > 0 {
			res = res + n - j
			value = int(s[i] - 'a')
			arr[value]--
			i++
		}
	}
	return res
}

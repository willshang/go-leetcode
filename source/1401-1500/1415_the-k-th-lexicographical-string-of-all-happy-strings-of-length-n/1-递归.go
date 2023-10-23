package main

func main() {

}

// leetcode1415_长度为n的开心字符串中字典序第k小的字符串
var arr []string

func getHappyString(n int, k int) string {
	arr = make([]string, 0)
	dfs(n, "")
	if len(arr) < k {
		return ""
	}
	return arr[k-1]
}

func dfs(n int, str string) {
	if len(str) > 1 && str[len(str)-1] == str[len(str)-2] {
		return
	}
	if len(str) == n {
		arr = append(arr, str)
		return
	}
	for i := 0; i < 3; i++ {
		char := 'a' + i
		dfs(n, str+string(char))
	}
}

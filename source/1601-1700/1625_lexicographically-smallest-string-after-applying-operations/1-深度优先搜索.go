package main

func main() {

}

var m map[string]bool
var res string

func findLexSmallestString(s string, a int, b int) string {
	res = s
	m = make(map[string]bool)
	dfs(s, a, b)
	return res
}

func dfs(s string, a, b int) {
	if m[s] == true {
		return
	}
	m[s] = true
	if s < res {
		res = s
	}
	dfs(add(s, a), a, b)
	dfs(s[b:]+s[:b], a, b)
}

func add(s string, a int) string {
	res := []byte(s)
	for i := 1; i < len(s); i = i + 2 {
		res[i] = byte('0' + (int(s[i]-'0')+a)%10)
	}
	return string(res)
}

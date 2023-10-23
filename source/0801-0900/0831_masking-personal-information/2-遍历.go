package main

import "strings"

func main() {

}

// leetcode831_隐藏个人信息
func maskPII(S string) string {
	if strings.Contains(S, "@") {
		S = strings.ToLower(S)
		arr := strings.Split(S, "@")
		return arr[0][:1] + "*****" + arr[0][len(arr[0])-1:] + "@" + arr[1]
	}
	res := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if '0' <= S[i] && S[i] <= '9' {
			res = append(res, S[i])
		}
	}
	n := len(res)
	str := "***-***-" + string(res[n-4:])
	if n > 10 {
		return "+" + strings.Repeat("*", n-10) + "-" + str
	}
	return str
}

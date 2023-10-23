package main

import "strings"

func main() {

}

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
	if len(res) == 10 {
		return "***-***-" + string(res[len(res)-4:])
	} else if len(res) == 11 {
		return "+*-***-***-" + string(res[len(res)-4:])
	} else if len(res) == 12 {
		return "+**-***-***-" + string(res[len(res)-4:])
	} else if len(res) == 13 {
		return "+***-***-***-" + string(res[len(res)-4:])
	}
	return string(res)
}

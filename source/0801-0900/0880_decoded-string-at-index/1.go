package main

func main() {

}

func decodeAtIndex(S string, K int) string {
	count := 0
	for i := 0; i < len(S); i++ {
		if '0' <= S[i] && S[i] <= '9' {
			value := int(S[i] - '0')
			count = count * value
		} else {
			count++
		}
	}
	for i := len(S) - 1; i >= 0; i-- {
		K = K % count

	}
	return ""
}

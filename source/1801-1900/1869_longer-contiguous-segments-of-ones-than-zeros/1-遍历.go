package main

func main() {

}

func checkZeroOnes(s string) bool {
	a, b := 0, 0
	prev := uint8(' ')
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == prev {
			count++
		} else {
			if prev == '0' {
				a = max(a, count)
			} else if prev == '1' {
				b = max(b, count)
			}
			count = 1
		}
		prev = s[i]
	}
	if prev == '0' {
		a = max(a, count)
	} else if prev == '1' {
		b = max(b, count)
	}
	return b > a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

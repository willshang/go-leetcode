package main

func main() {

}

// leetcode848_字母移位
func shiftingLetters(S string, shifts []int) string {
	arr := []byte(S)
	shifts = append(shifts, 0)
	for i := len(S) - 1; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
		arr[i] = 'a' + (S[i]-'a'+byte(shifts[i]))%26
	}
	return string(arr)
}

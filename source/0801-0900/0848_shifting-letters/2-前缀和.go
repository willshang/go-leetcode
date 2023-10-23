package main

func main() {

}

func shiftingLetters(S string, shifts []int) string {
	sum := 0
	for i := 0; i < len(shifts); i++ {
		sum = (sum + shifts[i]) % 26
	}
	arr := []byte(S)
	for i := 0; i < len(S); i++ {
		count := int(S[i] - 'a')
		arr[i] = byte((count+sum)%26 + 'a')
		sum = ((sum-shifts[i])%26 + 26) % 26
	}
	return string(arr)
}

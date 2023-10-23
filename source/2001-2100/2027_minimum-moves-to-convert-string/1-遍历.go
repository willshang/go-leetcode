package main

func main() {

}

func minimumMoves(s string) int {
	arr := []byte(s)
	res := 0
	for i := 0; i < len(s); i++ {
		if arr[i] == 'X' {
			count := 0
			for j := i; j < len(s); j++ {
				arr[j] = 'O'
				count++
				if count == 3 {
					break
				}
			}
			res++
		}
	}
	return res
}

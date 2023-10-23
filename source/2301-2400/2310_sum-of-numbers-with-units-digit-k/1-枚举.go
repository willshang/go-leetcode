package main

func main() {

}

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	if k == 0 {
		if num%10 == 0 {
			return 1
		}
		return -1
	}
	if num%2 == 1 && k%2 == 0 {
		return -1
	}
	for i := 1; i <= num; i++ {
		sum := i * k
		left := num - sum
		if left%10 == 0 {
			return i
		}
		if left < 0 {
			return -1
		}
	}
	return -1
}

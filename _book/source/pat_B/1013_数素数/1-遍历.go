package main

import "fmt"

func main() {
	var N, M int
	var num = 2
	var count int
	fmt.Scanf("%d %d", &M, &N)
	result := make([]int, 0)
	for {
		if count < N {
			if isPrime(num) {
				count++
				if count >= M {
					result = append(result, num)
				}
			}
			num++
		} else {
			break
		}
	}

	for i := 0; i < len(result); i++ {
		if i%10 != 0 {
			fmt.Print(" ")
		}
		fmt.Print(result[i])
		if i%10 == 9 {
			fmt.Println()
		}
	}
}

func isPrime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

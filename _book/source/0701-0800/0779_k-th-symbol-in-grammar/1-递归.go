package main

import "fmt"

func main() {
	fmt.Println(kthGrammar(4, 5))
}

func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	// N行K的数是由N-1行(K+1)/2的数来的
	temp := kthGrammar(N-1, (K+1)/2)
	if K%2 == 1 {
		return temp
	}
	return 1 - temp
}

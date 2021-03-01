package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}

func threeSumMulti(arr []int, target int) int {
	res := 0
	countArr := make([]int, 101)
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
	}
	for i := 0; i <= 100; i++ {
		// i < j < k
		for j := i + 1; j <= 100; j++ {
			k := target - i - j
			if j < k && k <= 100 {
				res = (res + countArr[i]*countArr[j]*countArr[k]) % 1000000007
			}
		}
		// i == j < k
		k := target - 2*i
		if i < k && k <= 100 {
			res = (res + countArr[i]*(countArr[i]-1)/2*countArr[k]) % 1000000007
		}
		// i < j == k
		if (target-i)%2 == 0 {
			j := (target - i) / 2
			if i < j && j <= 100 {
				res = (res + countArr[i]*countArr[j]*(countArr[j]-1)/2) % 1000000007
			}
		}
	}
	// i==j==k
	if target%3 == 0 {
		i := target / 3
		if 0 <= i && i <= 100 {
			res = (res + countArr[i]*(countArr[i]-1)*(countArr[i]-2)/6) % 1000000007
		}
	}
	return res % 1000000007
}

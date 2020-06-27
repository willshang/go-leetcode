package main

import "fmt"

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Println(largestTimeFromDigits([]int{0, 0, 0, 0}))
}

func largestTimeFromDigits(A []int) string {
	res := ""
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 4; l++ {
					if i != j && i != k && i != l &&
						j != k && j != l && k != l {
						hour := A[i]*10 + A[j]
						minute := A[k]*10 + A[l]
						if hour <= 23 && minute <= 59 {
							ans := fmt.Sprintf("%02d:%02d", hour, minute)
							if ans > res && res != "" {
								res = ans
							} else if res == "" {
								res = ans
							}
						}
					}
				}
			}
		}
	}
	return res
}

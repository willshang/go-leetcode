package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var A1, A2, A5 = 0, 0, 0
	var A4 = 0
	mapArr := make(map[int][]int)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		mapArr[num%5] = append(mapArr[num%5], num)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(mapArr[i]); j++ {
			if i == 0 && mapArr[i][j]%2 == 0 {
				A1 += mapArr[i][j]
			}
			if i == 1 && j%2 == 0 {
				A2 += mapArr[i][j]
			}
			if i == 1 && j%2 == 1 {
				A2 -= mapArr[i][j]
			}
			if i == 3 {
				A4 += mapArr[i][j]
			}
			if i == 4 && mapArr[i][j] > A5 {
				A5 = mapArr[i][j]
			}
		}
	}
	for i := 0; i < 5; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		if i == 0 && A1 == 0 || i != 0 && len(mapArr[i]) == 0 {
			fmt.Print("N")
			continue
		}
		if i == 0 {
			fmt.Print(A1)
		} else if i == 1 {
			fmt.Print(A2)
		} else if i == 2 {
			fmt.Print(len(mapArr[2]))
		} else if i == 3 {
			fmt.Print(fmt.Sprintf("%.1f", float64(A4)/float64(len(mapArr[i]))))
		} else if i == 4 {
			fmt.Print(A5)
		}
	}
}

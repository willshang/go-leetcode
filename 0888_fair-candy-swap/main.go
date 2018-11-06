package main

import "fmt"

func main()  {
	A := []int{1,2,5}
	B := []int{2,4}

	fmt.Println(fairCandySwap(A,B))
}


func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	Acoll := make(map[int]bool,len(A))

	for _, v := range A{
		sumA = sumA + v
		Acoll[v] = true
	}
	for _, value := range B{
		sumB = sumB + value
	}
	half := (sumA - sumB)/2
	a,b := 0,0
	for _,b = range B{
		a = b + half
		if Acoll[a] == true{
			break
		}
	}
	return []int{a,b}
}
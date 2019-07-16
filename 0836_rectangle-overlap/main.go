package main

import "fmt"

func main() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[1] < rec2[3] && rec1[0] < rec2[2] && rec2[1] < rec1[3] && rec2[0] < rec1[2] {
		return true
	}
	return false
}

/*func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	h1, v1 := cal(rec1)
	h2, v2 := cal(rec2)
	return isOverlap(h1,h2) && isOverlap(v1,v2)
}

func cal(rec []int)(h, v []int)  {
	h = []int{rec[0],rec[2]}
	v = []int{rec[1],rec[3]}
	return
}

func isOverlap(a,b []int)bool  {
	return !(a[1] <= b[0] || b[1] <= a[0])
}*/

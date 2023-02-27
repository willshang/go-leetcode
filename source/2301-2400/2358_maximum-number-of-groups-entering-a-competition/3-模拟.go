package main

func main() {

}

func maximumGroups(grades []int) int {
	n := len(grades)
	res := 0
	start := 1
	for start <= n {
		res++
		n = n - start
		start++
	}
	return res
}

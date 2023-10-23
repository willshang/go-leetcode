package main

func main() {

}

func countOperations(num1 int, num2 int) int {
	res := 0
	for num1 > 0 {
		res = res + num2/num1
		num1, num2 = num2%num1, num1
	}
	return res
}

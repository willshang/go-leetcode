package main

func main() {

}

// leetcode2011_执行操作后的变量值
func finalValueAfterOperations(operations []string) int {
	res := 0
	for i := 0; i < len(operations); i++ {
		if operations[i][1] == '+' {
			res++
		} else {
			res--
		}
	}
	return res
}

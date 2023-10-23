package main

func main() {

}

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 { // 转换为3进制，n*3^x，其中n为0或者1，不会为2
			return false
		}
		n = n / 3
	}
	return true
}

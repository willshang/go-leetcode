package main

func main() {

}

func minimumOneBitOperations(n int) int {
	// 在一组数的编码中，若任意两个相邻的代码只有一位二进制数不同
	// 则称这种编码为格雷码（Gray Code）
	res := 0
	for n > 0 {
		res = res ^ n
		n = n / 2
	}
	return res
}

package main

func main() {

}

func countEven(num int) int {
	a := num / 10 * 5 // 每10包含5个数
	b := 0            // 十位以上的数字之和
	for temp := num / 10; temp > 0; temp = temp / 10 {
		b = b + temp%10
	}
	if b%2 == 0 { // 偶数
		return a + (num%10+2)/2 - 1
	}
	return a + (num%10+1)/2 - 1
}

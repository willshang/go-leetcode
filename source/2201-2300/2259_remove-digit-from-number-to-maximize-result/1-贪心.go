package main

func main() {

}

func removeDigit(number string, digit byte) string {
	res := -1
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			res = i
			if i < len(number)-1 && number[i] < number[i+1] { // 贪心：找到第一个大于后面相邻的数
				break
			}
		}
	}
	return number[:res] + number[res+1:]
}

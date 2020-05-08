package main

import "fmt"

func main() {
	fmt.Println(addStrings("1234444", "2222229999999"))
}

func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	n1, n2 := len(num1), len(num2)
	a1, a2 := []byte(num1), []byte(num2)
	a1 = reverse(a1)
	a2 = reverse(a2)

	carry := 0
	buf := make([]byte, 0)
	for i := 0; i < n2; i++ {
		temp := 0
		if i < n1 {
			temp = int(a1[i] - '0')
		}
		temp = int(a2[i]-'0') + temp + carry
		if temp > 9 {
			buf = append(buf, byte(temp-10+'0'))
			carry = 1
		} else {
			buf = append(buf, byte(temp+'0'))
			carry = 0
		}
	}
	if carry == 1 {
		buf = append(buf, byte('1'))
	}
	return string(reverse(buf))
}

func reverse(arr []byte) []byte {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

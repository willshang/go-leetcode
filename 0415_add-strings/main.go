package main

import "fmt"

func main() {
	fmt.Println(addStrings("1234444","2222229999999"))
}
func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2){
		num1, num2 = num2, num1
	}

	n1, n2 := len(num1), len(num2)
	a1, a2 := []byte(num1), []byte(num2)

	carry := byte(0)

	buf := make([]byte,n2+1)
	buf[0] = '1'

	for i := 1; i <= n2;i++{
		if i <= n1{
			buf[n2+1-i] = a1[n1-i] - '0'
		}
		buf[n2+1-i] = buf[n2+1-i] + a2[n2-i] + carry

		if buf[n2+1-i] > '9'{
			buf[n2+1-i] = buf[n2+1-i] - 10
			carry = byte(1)
		}else {
			carry = byte(0)
		}
		fmt.Println(string(buf))
	}

	if carry == 1{
		return string(buf)
	}

	return string(buf[1:])
}
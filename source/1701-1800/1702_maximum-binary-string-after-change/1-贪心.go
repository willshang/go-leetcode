package main

func main() {

}

// leetcode1702_修改后的最大二进制字符串
func maximumBinaryString(binary string) string {
	flag := true
	rightOne := 0 // 记录第1个0后面1的数量
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			flag = false
		} else {
			if flag == false {
				rightOne++
			}
		}
	}
	if flag == true { // 全是1，直接返回
		return binary
	}
	// 首先：第1个0之前的1不需要移动。
	// 然后：把第1个0之后的1移到后面，后移：10=>01。
	// 最后：然后把中间的000，都变成110，00=>10。
	arr := make([]byte, len(binary))
	for i := 0; i < len(arr); i++ {
		arr[i] = '1'
	}
	arr[len(arr)-1-rightOne] = '0'
	return string(arr)
}

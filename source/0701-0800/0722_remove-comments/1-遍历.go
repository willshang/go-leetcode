package main

func main() {

}

// leetcode722_删除注释
func removeComments(source []string) []string {
	res := make([]string, 0)
	flag := false // 判断是否是块注释
	temp := make([]byte, 0)
	for i := 0; i < len(source); i++ {
		str := source[i]
		j := 0
		for j < len(str) {
			if flag == false && j+1 < len(str) && str[j] == '/' && str[j+1] == '*' {
				flag = true
				j = j + 2
				continue
			}
			if flag == true && j+1 < len(str) && str[j] == '*' && str[j+1] == '/' {
				flag = false
				j = j + 2
				continue
			}
			if flag == false && j+1 < len(str) && str[j] == '/' && str[j+1] == '/' {
				break
			}
			if flag == false {
				temp = append(temp, str[j])
			}
			j++
		}
		if flag == false && len(temp) > 0 {
			res = append(res, string(temp))
			temp = make([]byte, 0)
		}
	}
	return res
}

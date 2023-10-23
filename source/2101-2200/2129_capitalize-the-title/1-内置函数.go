package main

import "strings"

func main() {

}

// leetcode2129_将标题首字母大写
func capitalizeTitle(title string) string {
	arr := strings.Fields(title)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
		if len(arr[i]) > 2 {
			arr[i] = strings.Title(arr[i])
		}
	}
	return strings.Join(arr, " ")
}

package main

func main() {

}

// leetcode2273_移除字母异位词后的结果数组
func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	for i := 1; i < len(words); i++ {
		if judge(res[len(res)-1], words[i]) == false {
			res = append(res, words[i])
		}
	}
	return res
}

func judge(a, b string) bool {
	arr := [256]int{}
	for _, v := range a {
		arr[v]++
	}
	for _, v := range b {
		arr[v]--
	}
	return arr == [256]int{}
}

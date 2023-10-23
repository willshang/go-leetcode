package main

func main() {

}

// leetcode2194_Excel表中某个范围内的单元格
func cellsInRange(s string) []string {
	res := make([]string, 0)
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			res = append(res, string(i)+string(j))
		}
	}
	return res
}

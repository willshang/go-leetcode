package main

func main() {

}

func findLUSlength(strs []string) int {
	m := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		total := 1 << (len(strs[i]))
		for j := 0; j < total; j++ {
			s := ""
			for k := 0; k < len(strs[i]); k++ {
				if (j>>k)&1 != 0 {
					s = s + string(strs[i][k])
				}
			}
			m[s]++
		}
	}
	res := -1
	for k, v := range m {
		if v == 1 && len(k) > res {
			res = len(k)
		}
	}
	return res
}

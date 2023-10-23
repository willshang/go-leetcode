package main

func main() {

}

func longestWPI(hours []int) int {
	m := make(map[int]int)
	count := 0
	res := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			count++
		} else {
			count--
		}
		if count > 0 {
			res = i + 1
		} else {
			if _, ok := m[count]; ok == false {
				m[count] = i
			}
			// (count-(count-1))=1>0
			if _, ok := m[count-1]; ok == true {
				res = max(res, i-m[count-1])
			}

		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

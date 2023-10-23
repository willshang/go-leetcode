package main

func main() {

}

func findEvenNumbers(digits []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(digits); i++ {
		m[digits[i]]++
	}
	for i := 100; i < 1000; i = i + 2 {
		temp := make(map[int]int)
		value := i
		flag := true
		for value > 0 {
			temp[value%10]++
			if m[value%10] < temp[value%10] {
				flag = false
				break
			}
			value = value / 10
		}
		if flag == true {
			res = append(res, i)
		}
	}
	return res
}

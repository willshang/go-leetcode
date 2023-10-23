package main

func main() {

}

// leetcode1700_无法吃午餐的学生数量
func countStudents(students []int, sandwiches []int) int {
	a, b := 0, 0
	for i := 0; i < len(students); i++ {
		if students[i] == 0 {
			a++
		} else {
			b++
		}
	}
	for i := 0; i < len(sandwiches); i++ {
		if sandwiches[i] == 0 && a > 0 {
			a--
		} else if sandwiches[i] == 1 && b > 0 {
			b--
		} else {
			break
		}
	}
	return a + b
}

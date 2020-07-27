package main

import "fmt"

func main() {
	//fmt.Println(isLongPressedName("alex", "aaleex"))
	//fmt.Println(isLongPressedName("plpkoh", "plppkkh"))
	//fmt.Println(isLongPressedName("pyplrz", "ppyypllr"))
	fmt.Println(isLongPressedName("dfuyalc", "fuuyallc"))
	//fmt.Println(isLongPressedName("alex", "alexxr"))
	//fmt.Println(isLongPressedName("vtkgn", "vttkgnnggg"))
	//fmt.Println(isLongPressedName("vtkgn", "vttkgnngg"))
}

func isLongPressedName(name string, typed string) bool {
	i := 1
	j := 1
	countI := 0
	countJ := 0
	for i < len(name) || j < len(typed) {
		for i < len(name) && name[i] == name[i-1] {
			i++
			countI++
		}
		for j < len(typed) && typed[j] == typed[j-1] {
			j++
			countJ++
		}
		if name[i-1] != typed[j-1] || countJ < countI {
			return false
		}
		i++
		j++
		countI = 0
		countJ = 0
	}
	return name[len(name)-1] == typed[len(typed)-1]
}

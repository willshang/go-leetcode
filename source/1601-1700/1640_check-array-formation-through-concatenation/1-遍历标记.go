package main

func main() {

}

func canFormArray(arr []int, pieces [][]int) bool {
	for i := 0; i < len(pieces); i++ {
		value := pieces[i]
		length := len(value)
		flag := false
		for j := 0; j < len(arr); j++ {
			if arr[j] == value[0] {
				flag = true
				for k := j; k < j+length && k < len(arr); k++ {
					if arr[k] != value[k-j] {
						return false
					} else {
						arr[k] = 0
					}
				}
				break
			}
		}
		if flag == false {
			return false
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

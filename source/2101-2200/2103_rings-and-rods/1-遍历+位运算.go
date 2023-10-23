package main

func main() {

}

// leetcode2103_环和杆
func countPoints(rings string) int {
	m := map[byte]int{
		'R': 1,
		'G': 2,
		'B': 4,
	}
	arr := make([]int, 10)
	for i := 0; i < len(rings); i = i + 2 {
		v := int(rings[i+1] - '0')
		arr[v] = arr[v] | m[rings[i]]
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 7 {
			res++
		}
	}
	return res
}

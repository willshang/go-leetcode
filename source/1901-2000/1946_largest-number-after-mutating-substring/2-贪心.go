package main

func main() {

}

func maximumNumber(num string, change []int) string {
	res := []byte(num)
	for i := 0; i < len(num); i++ {
		if getValue(num[i]) < change[getValue(num[i])] {
			for i < len(num) && change[getValue(num[i])] >= getValue(num[i]) {
				res[i] = byte(change[getValue(num[i])] + '0')
				i++
			}
			break
		}
	}
	return string(res)
}

func getValue(c byte) int {
	return int(c - '0')
}

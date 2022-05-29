package main

func main() {

}

func largestGoodInteger(num string) string {
	res := ""
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i-1] == num[i-2] && num[i-2:i+1] > res {
			res = num[i-2 : i+1]
		}
	}
	return res
}

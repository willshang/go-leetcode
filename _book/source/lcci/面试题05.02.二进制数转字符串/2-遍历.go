package main

import "fmt"

func main() {
	fmt.Println(printBin(0.625))
}

func printBin(num float64) string {
	res := "0."
	value := float64(1)
	for i := 1; i <= 32; i++ {
		value = value / 2
		if num < value {
			res = res + "0"
			continue
		}
		res = res + "1"
		num = num - value
		if num == 0 {
			return res
		}
	}
	return "ERROR"
}

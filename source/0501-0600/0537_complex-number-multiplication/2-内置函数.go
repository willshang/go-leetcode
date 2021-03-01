package main

import (
	"fmt"
)

func main() {

}

func complexNumberMultiply(a string, b string) string {
	var a1, a2, b1, b2 int
	fmt.Sscanf(a, "%d+%di", &a1, &a2)
	fmt.Sscanf(b, "%d+%di", &b1, &b2)
	return fmt.Sprintf("%d+%di", a1*b1-a2*b2, a1*b2+a2*b1)
}

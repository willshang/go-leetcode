package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
func checkRecord(s string) bool {
	return !(strings.Count(s,"A") > 1 || strings.Contains(s,"LLL"))
}
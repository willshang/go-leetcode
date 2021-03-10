package main

import "strings"

func main() {

}

func checkOnesSegment(s string) bool {
	if strings.Contains(s, "01") {
		return false
	}
	return true
}

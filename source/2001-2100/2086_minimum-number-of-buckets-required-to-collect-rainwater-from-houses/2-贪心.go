package main

import "strings"

func main() {

}

func minimumBuckets(street string) int {
	if street == "H" || strings.Contains(street, "HHH") ||
		strings.HasSuffix(street, "HH") || strings.HasPrefix(street, "HH") {
		return -1
	}
	return strings.Count(street, "H") - strings.Count(street, "H.H")
}

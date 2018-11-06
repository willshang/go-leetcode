package main

import "fmt"

func main() {
	str := "abca"
	fmt.Println(validPalindrome(str))
}
func validPalindrome(s string) bool {
	i := 0
	j := len(s)-1
	for i < j{
		if s[i] != s[j]{
			return isPalindrome(s,i,j-1) || isPalindrome(s,i+1,j)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int)  bool{
	for i < j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

//func validPalindrome(s string) bool {
//	return helper([]byte(s), 0, len(s)-1,false)
//}
//
//func helper(bs []byte, low, high int,hasDeleted bool) bool  {
//	for low < high{
//		if bs[low] != bs[high]{
//			if hasDeleted{
//				return false
//			}
//			return helper(bs,low+1,high,true) ||
//				helper(bs,low,high-1,true)
//		}
//		low++
//		high--
//	}
//
//	return true
//}
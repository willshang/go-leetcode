package main

func main() {

}

// leetcode1328_破坏回文串
func breakPalindrome(palindrome string) string {
	if len(palindrome) < 2 {
		return ""
	}
	for i := 0; i < len(palindrome)/2; i++ {
		if palindrome[i] != 'a' {
			return palindrome[:i] + "a" + palindrome[i+1:]
		}
	}
	return palindrome[:len(palindrome)-1] + "b"
}

package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	} else if checkPolindrom(s) {
		return s
	}
	for lenght := len(s); lenght > 0; lenght-- {
		for start := 0; start <= len(s)-lenght; start++ {
			if checkPolindrom(s[start : start+lenght]) {
				return s[start : start+lenght]
			}
		}
	}
	return ""
}

func checkPolindrom(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func main() {

	st := "abb"

	fmt.Println(longestPalindrome(st))

}

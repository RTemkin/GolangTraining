package main

import "fmt"

func lenOfLongestSubstrbng(s string) int {

	var left, right, result int
	store := make(map[uint8]int)

	for right < len(s) {
		r := s[right]
		store[r] += 1
		for store[r] > 1 {
			l := s[left]
			store[l] -= 1
			left++
		}

		result = max(result, right-left+1)

		right++
	}
	return result
}

func main() {

	st := "abcabcddferrt"
	fmt.Println(lenOfLongestSubstrbng(st))

}

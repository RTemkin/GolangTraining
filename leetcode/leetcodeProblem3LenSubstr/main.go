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

func lenOfLongestSubstrbng2(st string) int {

	var left, right, result int

	mapHesh := make(map[uint8]int)

	for right < len(st) {
		r := st[right]
		mapHesh[r] += 1
		for mapHesh[r] > 1 {
			l := st[left]
			mapHesh[l] -= 1
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

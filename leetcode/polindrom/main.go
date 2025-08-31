package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	xStr := strconv.Itoa(x)

	if x < 0 {
		return false
	} else if x >= 0 && x < 10 {
		return true
	}

	i, j := 0, len(xStr)-1

	for i < len(xStr)/2 {
		if xStr[i] == xStr[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func main() {

	x := 12344321

	fmt.Println(isPalindrome(x))

}

package main

import "fmt"

func romanToInt(s string) int {

	var romanMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0

	for i, v := range s {

		val := romanMap[string(v)]
		result = result + val
		if i != 0 && romanMap[string(s[i])] > romanMap[string(s[i-1])] {
			result = result - 2*romanMap[string(s[i-1])]
		}
	}

	return result
}

func intToRoman(num int) string {

	numers := []int{}

	for i := 1; num != 0; {
		numers = append(numers, num%10*i)
		num = (num - num%10) / 10
		i *= 10

	}

	fmt.Println(numers)
	return "one"
}

func main() {

	st := "XXXIVII"

	fmt.Println(romanToInt(st))

	intToRoman(3099)

}

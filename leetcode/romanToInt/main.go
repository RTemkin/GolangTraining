package main

import (
	"fmt"
	"strings"
)

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

	intmap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	numers := []int{}

	for i := 1; num != 0; {
		numers = append(numers, num%10*i)
		num = (num - num%10) / 10
		i *= 10
	}

	var builder strings.Builder
	
	for i := len(numers) - 1; i >= 0; i-- {
		if numers[i]>1000{
			kof := numers[i]/1000
			builder.WriteString(intmap[numers[i]])
		}
		
		builder.WriteString(intmap[numers[i]])
	}

	return builder.String()
}

func main() {

	st := "XXXIVII"

	fmt.Println(romanToInt(st))

	fmt.Println(intToRoman(1964))

}

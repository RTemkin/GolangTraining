package main

import "fmt"

func interSelection(sl1, sl2 []int) []int {

	contMap := make(map[int]int)
	result := []int{}

	for _, val := range sl1 {
		if _, ok := contMap[val]; !ok {
			contMap[val] = 1
		} else {
			contMap[val] += 1
		}
	}

	for _, val := range sl2 {
		if couter, ok := contMap[val]; ok && couter > 0 {
			contMap[val] -= 1
			result = append(result, val)
		}
	}

	return result

}

func main() {

	a := []int{23, 3, 1, 2, 23}
	b := []int{6, 2, 4, 23, 23, 23}

	fmt.Printf("%v\n", interSelection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}

	fmt.Printf("%v\n", interSelection(a, b))

}

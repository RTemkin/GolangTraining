package main

import (
	"fmt"
	"math/rand"
)
func GeneratorDubl(n int) []int {

	dublicat := make([]int, n, n)
	for i := 0; i < n; i++ {
		dublicat[i] = rand.Intn(31)
	}
	return dublicat
}

func Delldublicat(sl []int) []int {
	mapVal := make(map[int]bool)
	res := []int{}

	for _, val := range sl {
		if _, ok := mapVal[val]; !ok {
			mapVal[val] = true
			res = append(res, val)
		}
	}
	return res
}

func SortBabls(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}
	for i := 0; i < len(sl)-1; i++ {
		j := i + 1
		for j < len(sl) {
			if sl[j] < sl[i] {
				sl[i], sl[j] = sl[j], sl[i]
			}
			j++
		}
	}
	return sl
}

func main() {
	dubls := GeneratorDubl(1000)

	// fmt.Println(dubls)

	res := Delldublicat(dubls)

	fmt.Println(res)

	// res2 := []int{11, 2, 3, 3, 2, 5, 12, 11, 17, 4, 7, 6, 18, 5, 16, 1, 0, 4, 8}
	fmt.Println(SortBabls(res))

}

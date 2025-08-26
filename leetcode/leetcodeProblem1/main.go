package main

import "fmt"

func sum(nums []int, target int) []int {

	resMap := make(map[int]int)

	for i, val := range nums {
		comp := target - val

		if idx, ok := resMap[comp]; ok {
			return []int{i, idx}
		}

		resMap[val] = i
	}

	return nil
}

func sum2(nums []int, target int) []int {
	res := make([]int, 2)

	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		for j < len(nums) {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
			j++
		}
	}
	return nil
}

func sum3(nums []int, target int) []int {
	mapss := make(map[int]int)

	for i, val := range nums {
		comp := target - val
		if indx, ok := mapss[comp]; ok {
			return []int{i, indx}
		}
		mapss[val] = i
	}

	return nil
}

func sum4(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++{
		j := i+1
		for j < len(nums){
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
			j++
		}
	}
	return nil
}

func main() {

	nums := []int{1, 4, 5, 6, 7, 8}
	target := 12

	fmt.Println(sum3(nums, target))

	mapa := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	slice := []string{"qwe", "q", "qqq", "aaa", "three"}

	for i, val := range slice {

		if idx, ok := mapa[val]; ok {
			fmt.Println(i, idx, ok, val)
		} else {
			fmt.Println("Noooooo", i, idx, ok, val)
		}
	}

	for _, val := range mapa {

		fmt.Println(val)

	}

}

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Вариант 1 горизонтальный просмотр
	// 	prefix := strs[0]

	// 	for i := 1; i < len(strs); i++ {
	// 		for strings.Index(strs[i], prefix) != 0 {
	// 			prefix = prefix[:len(prefix)-1]
	// 			if prefix == "" {
	// 				return ""
	// 			}
	// 		}
	// 	}

	// 	return prefix

	//Вариант 2 вертикальный просмотр
	// for i := 0; i < len(strs[0]); i++ {
	// 	c := strs[0][i]
	// 	for j := 1; j < len(strs); j++ {
	// 		if i == len(strs[j]) || strs[j][i] != c {
	// 			return strs[0][:i]
	// 		}
	// 	}
	// }
	// return strs[0]



	first := strs[0]

	for _, st := range strs{
		if len(st)<len(first){
			first = st
		}
	}
	
	for i:=0; i<len(first);i++{
		for j:=0; j<len(strs);j++{
			second := strs[j]

			if first[i] != second[i]{
				return first[:i]
			}
		}
		
	}

	return first
}

func main() {

	strs := []string{"qwttytv", "qwe", "qwertyy", "qwrty"}
	fmt.Println(longestCommonPrefix(strs))

}

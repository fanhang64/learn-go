package main

import (
	"fmt"
	"strings"
)

func main(){
	hanziCount("abcdef123AAA!!@@")
	hanziCount("汉字的个数")
}


func hanziCount(src string)  {
	letters := "abcdefghijklmnopqrstuvwxyz"
	letters += strings.ToUpper(letters)

	nums := "0123456789"

	lettersCount := 0
	numsCount := 0
	chineseCount := 0

	for _, i := range src{
		switch {
		case strings.ContainsRune(letters, i) == true:
			lettersCount++
		case strings.ContainsRune(nums, i) == true:
			numsCount++
		default:
			chineseCount++
			
		}
	}
	fmt.Println(lettersCount, numsCount, chineseCount)
}
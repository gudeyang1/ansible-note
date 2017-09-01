package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "123vfs?>"
	letters := "abcdefghijklmnopqrstuvwxyz"
	letters = letters + strings.ToUpper(letters)
   	nums := "0123456789"
	num := 0
	aplap := 0
	other :=0
	for _,i := range a{
		switch  {
		case strings.ContainsRune(letters,i):
			aplap ++

		case strings.ContainsRune(nums,i):
			num ++
		default:
			other ++
		}

	}
		fmt.Println(aplap)
		fmt.Println(num)
		fmt.Println(other)

}

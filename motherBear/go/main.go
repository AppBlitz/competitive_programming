package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome(createSlice(strings.ToLower("Si nummi immunis"))))
}

func createSlice(word string) (result string) {
	slices := string([]rune(word))
	for v := 0; v < len(slices); v += 1 {
		if string(slices[v]) != "." && string(slices[v]) != "!" && string(slices[v]) != "," && string(slices[v]) != "?" {
			result += string(slices[v])
		}
	}
	return result
}

func isPalindrome(word string) bool {
	auxiliary := ""
	slices := string([]rune(word))
	for v := len(slices) - 1; v >= 0; v -= 1 {
		auxiliary += string(slices[v])
	}
	if auxiliary == word {
		return true
	} else {
		return false
	}
}

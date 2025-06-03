package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for range n {
		scanner.Scan()
		line := scanner.Text()
		number, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if err != nil || number < 0 {
			continue
		}
		iterations, palindrome := computePalindrome(number)
		fmt.Printf("%d %d\n", iterations, palindrome)
	}
}

func computePalindrome(num int64) (int, int64) {
	var iterations int
	for iterations < 1000 {
		if isPalindrome(num) {
			return iterations, num
		}
		num += reverseNumber(num)
		iterations++
	}
	return iterations, num // Por restricciones del problema, nunca se alcanza.
}

func reverseNumber(num int64) int64 {
	var reversed int64
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return reversed
}

func isPalindrome(num int64) bool {
	s := strconv.FormatInt(num, 10)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

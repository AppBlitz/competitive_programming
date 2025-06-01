package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	for {
		n, _ := fmt.Scanf("%d", &number)
		if number == 0 || n != 1 {
			break
		}
		if number <= 2000000000 {
			fmt.Printf("%d\n", verificationLength(number))
		}
	}
}

func verificationLength(number int) int {
	auxiliary := strconv.Itoa(sumDigits(number))
	n, _ := strconv.Atoi(auxiliary)
	if len(auxiliary) > 1 {
		return verificationLength(n)
	}
	return n
}

func sumDigits(number int) (sum int) {
	for number != 0 {
		sum = (number % 10) + sum
		number = number / 10
	}
	return sum
}

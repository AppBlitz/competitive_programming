package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	for number != "END" {
		_, error := fmt.Scanf("%s", &number)
		if n, erro := strconv.Atoi(number); erro == nil && error == nil && calculatingDigists(n) <= 1000000 {
			fmt.Printf("%d\n", verificationNumber(n, 0))
		}
	}
}

func verificationNumber(number, i int) int {
	if number == 1 {
		if calculatingDigists(number) == number {
			return i + 1
		} else {
			return 0
		}
	} else {
		return verificationNumber(calculatingDigists(number), i+1)
	}
}

func calculatingDigists(number int) (sumDigits int) {
	for number != 0 {
		sumDigits = sumDigits + 1
		number = number / 10
	}
	return sumDigits
}

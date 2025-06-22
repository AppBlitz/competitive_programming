package main

import (
	"fmt"
)

func main() {
	var test int
	var number int
	_, erro := fmt.Scanf("%d", &test)
	if erro == nil {
		for test >= 0 {
			_, ok := fmt.Scanf("%d", &number)
			if ok == nil {
				fmt.Printf("%d\n", multiplicationNumber(number))
			}
			test -= 1
		}
	}
}

func multiplicationNumber(number int) int {
	var result int
	var auxiliary int
	result = number * 567
	result /= 9
	result += 7492
	result *= 235
	result /= 47
	result -= 498
	for v := 1; v <= 2; v += 1 {
		auxiliary = result % 10
		result = result / 10
	}
	if auxiliary < 0 {
		auxiliary *= -1
	}
	return auxiliary
}

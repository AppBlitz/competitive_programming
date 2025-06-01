package main

import "fmt"

func main() {
	var i, j int
	for {
		_, err := fmt.Scan(&i, &j)
		if err != nil {
			break
		}
		fmt.Printf("%d %d %d\n", i, j, runNumber(i, j))
	}
}

func runNumber(i int, j int) (result int) {
	result = 0
	if i > j {
		for v := j; v <= i; v += 1 {
			auxliary := calculateNumber(v)
			if auxliary > result {
				result = auxliary
			}
		}
	} else {
		for v := i; v <= j; v += 1 {
			auxliary := calculateNumber(v)
			if auxliary > result {
				result = auxliary
			}
		}
	}
	return result
}

func calculateNumber(number int) (count int) {
	count = 1
	for number != 1 {
		switch number % 2 {
		case 1:
			{
				number = (number * 3) + 1
				count = count + 1
			}
		case 0:
			{
				number = number / 2
				count = count + 1
			}
		}
	}
	return count
}

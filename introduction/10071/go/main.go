package main

import "fmt"

func main() {
	var v int
	var t int
	for {
		_, erro := fmt.Scanf("%d\n %d", &v, &t)
		if erro != nil {
			break
		}
		if v >= -100 && v <= 100 && t <= 200 {
			fmt.Printf("%d\n", 2*v*t)
		}
	}
}

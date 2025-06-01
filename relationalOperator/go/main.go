package main

import "fmt"

func main() {
	var iterator int
	_, erro := fmt.Scanf("%d", &iterator)
	if erro == nil {
		for iterator != 0 {
			var a, b int
			nu, _ := fmt.Scanf("%d\n %d", &a, &b)
			if nu != 0 {
				if a <= 1000000001 && b <= 1000000001 {
					fmt.Printf("%s\n", verificationNumber(a, b))
				}
			}
			iterator = iterator - 1
		}
	}
}

func verificationNumber(a, b int) string {
	if a > b {
		return ">"
	} else if a < b {
		return "<"
	} else {
		return "="
	}
}

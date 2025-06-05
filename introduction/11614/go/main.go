package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64
	_, erro := fmt.Scanf("%d", &n)
	if erro == nil && n >= 0 && n < int64(math.Pow10(18)) {

		var warriors int64
		var i int64
		for i = 0; i < n; i += 1 {
			_, erro := fmt.Scanf("%d", &warriors)
			if erro == nil {
				fmt.Printf("%d\n", CalculatingRows(warriors))
			}
		}
	}
}

func CalculatingRows(warriors int64) (result int) {
	if warriors > 0 {
		auxiliary := int(math.Sqrt(float64((8 * warriors) + 1)))
		result = (auxiliary - 1) / 2
		return result
	} else {
		return 0
	}
}

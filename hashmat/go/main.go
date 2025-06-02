package main

import (
	"fmt"
	"math"
)

func main() {
	var soldierOne, soldierTwo int
	for {
		_, erro := fmt.Scanf("%d\n%d", &soldierOne, &soldierTwo)
		if erro != nil {
			break
		} else {
			if soldierOne <= int(math.Pow(2, 32)) && soldierTwo <= int(math.Pow(2, 32)) {
				fmt.Printf("%d\n", calculaDiference(soldierOne, soldierTwo))
			}
		}
	}
}

func calculaDiference(soldierOne, soldierTwo int) (diference int) {
	if soldierOne > soldierTwo {
		diference = soldierOne - soldierTwo
	} else {
		diference = soldierTwo - soldierOne
	}
	return diference
}

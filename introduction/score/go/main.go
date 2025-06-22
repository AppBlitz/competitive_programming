package main

import (
	f "fmt"
	s "strings"
)

func main() {
	var testCase int
	var stringCase string
	var sumString int
	_, erro := f.Scan(&testCase)
	if erro == nil {
		for i := 0; i < testCase; i += 1 {
			auxiliary := 1
			_, ok := f.Scan(&stringCase)
			if ok == nil && len(stringCase) > 0 && len(stringCase) <= 80 {
				value := s.TrimSpace(stringCase)
				for _, vs := range value {
					if vs == 79 {
						sumString += auxiliary
						auxiliary += 1
					} else {
						auxiliary = 1
					}
				}
				f.Printf("%d\n", sumString)
			}
		}
	}
}

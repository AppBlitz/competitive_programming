package main

import (
	"fmt"
	"strings"
)

func main() {
	mp := map[string]string{
		"HELLO":        "ENGLISH",
		"HOLA":         "SPANISH",
		"HALLO":        "GERMAN",
		"BONJOUR":      "FRENCH",
		"CIAO":         "ITALIAN",
		"ZDRAVSTVUJTE": "RUSSIAN",
	}
	var s string
	i := 1
	for s != "#" && i <= 2000 {
		_, erro := fmt.Scanf("%s", &s)
		if erro == nil && s != "#" && len(s) <= 14 {
			value, ok := mp[strings.ToUpper(s)]
			if ok {
				fmt.Printf("Case %d: %s\n", i, value)
			} else {
				fmt.Printf("Case %d: %s ", i, "UNKNOWN")
			}
		} else {
			break
		}
		i += 1
	}
}

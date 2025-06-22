package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(countWords(line))
	}
}

func countWords(line string) int {
	re := regexp.MustCompile(`[A-Za-z]+`)
	words := re.FindAllString(line, -1)
	return len(words)
}

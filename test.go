package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("Q: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Print("A: ")
		userInput := scanner.Text()

		// Remove "吗" and "?"
		userInput = strings.ReplaceAll(userInput, "吗", "")
		userInput = strings.ReplaceAll(userInput, "?", "")
		userInput = strings.ReplaceAll(userInput, "？", "")

		fmt.Println(userInput)
	}
}

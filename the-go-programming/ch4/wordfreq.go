package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your input: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading the input")
		os.Exit(1)
	}
	words := strings.Split(input, " ")

	for _, word := range words {
		fmt.Println(word)
	}
	fmt.Println("Total count is:", len(words))
}

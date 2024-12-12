package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func executeCharCounter() {
	fmt.Print("Please input: ")
	scanner := bufio.NewScanner(os.Stdin)

	if res := scanner.Scan(); !res {
		err := scanner.Err()
		if err == io.EOF {
			return
		} else {
			fmt.Printf("error: %v", err.Error())
		}
	}
	text := scanner.Text()
	counter := make(map[string]int)
	charCounter(text, counter)

	fmt.Printf("Counter: %v", counter)
}

func main() {
	// executeCharCounter()
	fmt.Println(executeCharCounterMultiInput())
}

func charCounter(text string, counter map[string]int) {

	for _, r := range text {
		s := string(r)
		if s == "" || s == " " {
			continue
		}
		if ok := counter[s]; ok != 0 {
			counter[s]++
		} else {
			counter[s] = 1
		}
	}
}

func executeCharCounterMultiInput() (map[string]int, error) {
	counter := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please input: ")
		if !scanner.Scan() {
			break
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			break
		}
		charCounter(text, counter)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	return counter, nil
}

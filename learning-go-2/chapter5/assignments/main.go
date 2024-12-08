package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var (
	add = func(i1, i2 int) int { return i1 + i2 }
	sub = func(i1, i2 int) int { return i1 - i2 }
	mul = func(i1, i2 int) int { return i1 * i2 }
	div = func(i1, i2 int) int { return i1 / i2 }
)

func main() {
	// executeCalculator()
	// executeFileLen()
	executePrefixer()
}
func executePrefixer() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
func executeFileLen() {
	length, err := fileLen("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Length:", length)
}
func executeCalculator() {
	pairSet := [][]string{{"1", "+", "3"}, {"2", "-", "5"}, {"15", "/", "5"}, {"4", "*", "2"}, {"4", "test", ""}, {"3", "+", ""}, {"abc", "", ""}}

	for _, pair := range pairSet {
		p1 := pair[0]
		op := pair[1]
		p2 := pair[2]
		result, err := simpleCalculator(p1, op, p2)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}
func simpleCalculator(s1, op, s2 string) (int, error) {
	switch op {
	case "+":
		s1, err1 := strconv.Atoi(s1)
		s2, err2 := strconv.Atoi(s2)
		if err1 != nil || err2 != nil {
			return 0, errors.New(fmt.Sprintf("cannot cast the given values: val1: %v val2: %v", s1, s2))
		}
		return add(s1, s2), nil
	case "-":
		s1, err1 := strconv.Atoi(s1)
		s2, err2 := strconv.Atoi(s2)
		if err1 != nil || err2 != nil {
			return 0, errors.New(fmt.Sprintf("cannot cast the given values: val1: %v val2: %v", s1, s2))
		}
		return sub(s1, s2), nil
	case "*":
		s1, err1 := strconv.Atoi(s1)
		s2, err2 := strconv.Atoi(s2)
		if err1 != nil || err2 != nil {
			return 0, errors.New(fmt.Sprintf("cannot cast the given values: val1: %v val2: %v", s1, s2))
		}
		return mul(s1, s2), nil
	case "/":
		s1, err1 := strconv.Atoi(s1)
		s2, err2 := strconv.Atoi(s2)
		if err1 != nil || err2 != nil {
			if s2 == 0 {
				return 0, errors.New("cannot divide by zero")
			}
		}
		return div(s1, s2), nil
	default:
		return 0, errors.New("unsupported operation")
	}
}

/*
*
The defer keyword in Go delays the execution of a function until the surrounding function exits. See "defer."
*
*/
func fileLen(filename string) (int, error) {
	path, err := os.Getwd()
	if err != nil {
		return 0, err
	}

	file, err := os.Open(filepath.Join(path, filename))
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	length := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				length += len([]byte(line))
				break
			}
			return length, err
		}
		length += len([]byte(line))
	}

	return length, nil
}

func prefixer(input string) func(s string) string {
	return func(s string) string { return fmt.Sprintf("%s %s", input, s) }
}

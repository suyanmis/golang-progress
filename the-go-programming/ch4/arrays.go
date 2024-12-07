package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	TRY
	RMB
	JPY
	RUB
)

func main() {
	fmt.Println(flag.Args())
	strategy := flag.String("s", "", "Strategy to use for: sha256, sha512, sha384")
	data := flag.String("d", "", "Data to encrypt")
	flag.Parse()
	if *strategy == "" || *data == "" {
		fmt.Println("Error: both -s (strategy) and -d (data) flags are required")
		flag.Usage()
		os.Exit(1)
	}

	shaCounter(*strategy, *data)

	// index and value pairing
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥", RUB: "rub"}
	for i, c := range symbol {
		fmt.Print(i, c, " ")
	}
	var q1 [3]int = [3]int{1, 2, 3}
	q2 := [3]int{1, 2, 3}
	q3 := [...]int{1, 2, 3}
	var q4 [3]int
	fmt.Println()
	fmt.Println(q1, q2, q3, q4)
	fmt.Println(q4 == [3]int{})

	var q5 [3]int = [3]int{1, 2, 3}
	// q5 = [4]int{1, 2, 3, 4} // Compilation error

	fmt.Println(q5)

	r := [...]int{9: -1} // assigns the index 9 to -1, rest is 0

	fmt.Println(r, len(r))

	sha256Comparer("test", "tesd")
}

func sha256Comparer(s1 string, s2 string) {
	sha1 := sha256.Sum256([]byte(s1))
	sha2 := sha256.Sum256([]byte(s2))
	diffCount := 0
	fmt.Printf("SHA1 and SHA2 lengths: %d, %d\n", len(sha1), len(sha2))
	for i := 0; i < sha256.Size; i++ {
		if sha1[i] != sha2[i] {
			diffCount++
		}
	}
	fmt.Printf("sha256 byte difference: %d", diffCount)
}

func shaCounter(strategy string, input string) {
	fmt.Printf("Strategy: %s, input: %s\n", strategy, input)

	switch strategy {
	case "sha256":
		output := sha256.Sum256([]byte(input))
		fmt.Printf("SHA256: %x\n", string(output[:]))
	case "sha384":
		output := sha512.Sum384([]byte(input))
		fmt.Printf("SHA384: %x\n", string(output[:]))
	case "sha512":
		output := sha512.Sum384([]byte(input))
		fmt.Printf("SHA256: %x\n", string(output[:]))
	}
}

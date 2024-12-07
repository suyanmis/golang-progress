package main

import "fmt"

func main() {

	samples := []string{"hello", "apple_π!"}

outer:
	for _, sample := range samples {
	inner:
		for _, r := range sample {
			fmt.Println("Rune: ", string(r))
			if r == 'e' {
				break inner
			}
			if r == 'p' {
				continue outer
			}
		}
	}

	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}

	// by default cases don't fall through
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is short!")
		case 5:
			fmt.Println(word, "is great!")
		default:
			fmt.Println(word, "is not short or great!")
		}
	}

loop:
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is short!")
			break // not as intended
		case 5:
			fmt.Println(word, "is great!")
			break // not as intended
		case 6, 7, 8:
			break loop // breaks the loop
		}
	}
}

package main

import (
	"fmt"
)

func main() {
TESTING:
	fmt.Println("HEY")

	for {
		goto TESTING
	}

}

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	c, ok := w.(*bytes.Buffer)
	if !ok {
		fmt.Println("Type assertion.")
	}
	fmt.Println(f, c)
}

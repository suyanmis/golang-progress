package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	fmt.Println()
	return len(p), nil
}

func main() {
	// text := []byte("hello")
	// bc := ByteCounter(0)
	// bc.Write(text)

	// fmt.Println(bc)
	// bc = 0
	// fmt.Fprintf(&bc, "%s, %s", text, "serhat")
	// fmt.Println(bc)

	m := MyReader{}
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// m.Read([]byte("h1:testing"))
	m.Read(data)
}

func (total *ByteCounter) Write2(p []byte) (int, error) {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return int(*total), err
	}
	defer f.Close()

	const chunkSize = 8096
	b := make([]byte, chunkSize)

	for {
		readTotal, err := f.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		*total += ByteCounter(readTotal)
		fmt.Println(string(b[:readTotal]))
	}

	return int(*total), nil

}

type MyReader struct {
}

type HtmlElement struct {
	Element      *string
	ChildElement *HtmlElement
	Content      *string
}

func (m MyReader) Read(p []byte) (n int, err error) {
	currWord := ""
	hEl := HtmlElement{}
	for i := range p {
		letter := strings.Trim(string(p[i]), "")
		if letter == ":" {
			// set element and start content after here
			hEl.Element = &currWord
			currWord = ""
		} else {
			currWord += letter
			fmt.Println(currWord)
		}
	}
	hEl.Content = &currWord
	data := htmlParser(hEl)
	os.WriteFile("sample.html", data, 0664)
	return 0, nil
}
func htmlParser(hEl HtmlElement) []byte {
	return []byte(fmt.Sprintf("<%s>%s</%s>", *hEl.Element, *hEl.Content, *hEl.Element))
}

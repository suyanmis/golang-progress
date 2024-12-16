package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) process() {
	fmt.Println("Processing")
}

type Logic interface {
	process()
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	c.L.process()
}

func main() {
	c := Client{
		L: LogicProvider{},
	}

	c.Program()
}

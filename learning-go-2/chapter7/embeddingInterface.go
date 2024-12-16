package main

type Reader interface {
	Read(p []byte) (n int, e error)
}

type Closer interface {
	Close() error
}

// Just as we can embed a concrete type in a struct we can embed interface in a struct as well.

type MyReader struct {
	Reader Reader
	buffer []byte
}
type ReaderCloser interface {
	Reader
	Closer
}

//If multiple implementations of an interface are passed into the function, this will mean creating multiple functions with repeated logic

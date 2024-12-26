package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type MyErr struct {
	message string
}

func (e *MyErr) Error() string {
	if len(e.message) != 0 {
		return e.message
	}
	return "Custom Error message"
}
func main() {
	// err := ReadFile()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	_, err := os.Open("not_exist.txt")
	if os.IsNotExist(err) {
		log.Println("File doesn't exist")
		fmt.Printf("%#v\n", err)
	}

}

// os.IsNotExist
func IsNotExist(er error) bool {
	if e, ok := er.(*os.PathError); ok {
		return e == fs.ErrNotExist
	}
	return false
}
func ReadFile() error {
	// return &MyErr{message: "Customized Message For Read File"}
	return &MyErr{}
}

// fs.PathError
type PathError struct {
	Op   string
	Path string
	err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.err.Error()
}

// func underlyingError(err error) error {
// 	switch err := err.(type) {
// 	case *PathError:
// 		return err.Err
// 	case *LinkError:
// 		return err.Err
// 	case *SyscallError:
// 		return err.Err
// 	}
// 	return err
// }

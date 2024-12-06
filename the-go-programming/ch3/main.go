package main

import "fmt"

func main() {
	fmt.Printf("\370")

	fmt.Println(isAnagram("test"))
	fmt.Println(isAnagram("tetet"))
	fmt.Println(isAnagram("exaxe"))
	fmt.Println(isAnagram("morrom"))
	fmt.Println(isAnagram("mordorom"))
}

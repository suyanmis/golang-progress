package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	ss := StringSlice{"test", "my", "string", "a", "b", "c", "d", "e", "f", "g", "h", "t"}
	sort.Sort(ss)
	sort.Strings(ss)
	// sort.IntSlice()
	// sort.Float64Slice()
	fmt.Printf("ss: %v\n", ss)
}

package main

import (
	"fmt"
	"time"
)

type Incrementer interface {
	Increment()
}
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdated: %v", c.total, c.lastUpdated)
}

// type Stringer interface {
// 	String() string
// }

func main() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter
	myStringer = valueCounter
	myIncrementer = pointerCounter
	// myIncrementer = valueCounter // this works when the implemented method does not use the pointer. Otherwise we should pass the address.
	myIncrementer = &valueCounter

}

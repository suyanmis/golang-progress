package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// flag.Value
	flag := flag.NewFlagSet("sleepFlagger", flag.ExitOnError)
	today := CelciusFlag{degree: 65.5, unit: "F"}
	flag.Var(&today, "unit", "temperature unit")
	// sleep := flag.Duration("period", 3*time.Second, "sleep period")
	if err := flag.Parse(os.Args[1:]); err != nil {
		fmt.Println(flag.Usage)
		os.Exit(1)
	}
	fmt.Println("Today:", today)
	// time.Sleep(*sleep)
	fmt.Println("Done")
}

// Interface to implement for new flag value:
// type Value interface {
// 	String() string
// 	Set(string) error
// }

type CelciusFlag struct {
	degree float64
	unit   string
}

func (c *CelciusFlag) String() string {
	return fmt.Sprintf("degree: %f unit: %s", c.degree, c.unit)
}
func (c *CelciusFlag) Set(s string) error {
	if !strings.Contains(s, "C") && !strings.Contains(s, "F") {
		return errors.New("incorrect unit")
	}

	if _, err := fmt.Sscanf(s, "%f%s", &c.degree, &c.unit); err != nil {
		return err
	}

	return nil
}

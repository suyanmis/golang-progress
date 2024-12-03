package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	unitConverter()
}

func unitConverter() {
	inputMessages := [4]string{
		"Available options: T(Temperature, Units ->(F,C)), W(Weight, Units ->(LB, KG)), D(Distance, Units -> (M,FT))",
		"Current Unit",
		"Target Unit",
		"Value to convert",
	}
	fs := flag.NewFlagSet("start", flag.ContinueOnError)
	options := fs.String("o", "", inputMessages[0])
	cu := fs.String("cu", "", inputMessages[1])
	tu := fs.String("tu", "", inputMessages[2])
	v := fs.String("v", "", inputMessages[3])
	err := fs.Parse(os.Args[1:])

	var inputs [4]string
	if err != nil || *options == "" || *cu == "" || *tu == "" || *v == "" {
		fmt.Println("Options have failed or are incomplete. We will accept the user input over command line.")
		for i := 0; i < 4; i++ {
			fmt.Printf("%s: ", inputMessages[i])
			var input string
			fmt.Scanln(&input) // Waits for user input
			inputs[i] = input
		}
	} else {
		inputs[0] = *options
		inputs[1] = *cu
		inputs[2] = *tu
		inputs[3] = *v
	}

	response := process(inputs)
	fmt.Println("Response:", response)
}

func process(inputs [4]string) string {
	unit := inputs[0]
	current := inputs[1]
	target := inputs[2]
	value := inputs[3]
	var response float64

	switch unit {
	case "T":
		response = handleTemperature(current, target, value)
	case "D":
		response = handleDistance(current, target, value)
	case "W":
		response = handleWeight(current, target, value)
	}

	return fmt.Sprintf("Conversion from %s to %s: %f", current, target, response)
}

func castStringToFloat(value string) float64 {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Printf("Error while parsing the value: %s. Error: %v\n", value, err)
		os.Exit(1)
	}
	return val
}

func handleWeight(current, target, value string) float64 {
	val := castStringToFloat(value)
	if current == "KG" && target == "LB" {
		val = convertKgToLb(val)
	} else if current == "LB" && target == "KG" {
		val = convertLbToKg(val)
	}
	return val
}

func convertKgToLb(val float64) float64 {
	return val * 2.2
}

func convertLbToKg(lb float64) float64 {
	return lb / 2.2
}

func handleDistance(current, target, value string) float64 {
	val := castStringToFloat(value)
	if current == "M" && target == "FT" {
		val = convertMeterToFeet(val)
	} else if current == "FT" && target == "M" {
		val = convertFeetToMeter(val)
	}
	return val
}

func convertMeterToFeet(meter float64) float64 {
	return meter * 3.28084
}

func convertFeetToMeter(f float64) float64 {
	return f * 0.3048
}

func handleTemperature(current string, target string, value string) float64 {
	val := castStringToFloat(value)
	if current == "C" && target == "F" {
		val = convertCtoF(val)
	} else if current == "F" && target == "C" {
		val = convertFtoC(val)
	}
	return val
}

func convertCtoF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func convertFtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// instead of using interface{} type new code uses any{} type
	// useful when there is an uncertainty.

	data := map[string]any{}
	cont, err := os.ReadFile("test.json")
	if err != nil {
		log.Fatal("Error: %s", err.Error())
	}
	json.Unmarshal(cont, &data)

	fmt.Println(data["name"], data["lastname"], data["age"])

}

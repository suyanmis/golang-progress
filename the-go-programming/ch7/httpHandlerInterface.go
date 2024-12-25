package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollar
type dollar float32

func (d dollar) String() string {
	return fmt.Sprintf("$%2.f", d)
}
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		err := fmt.Sprintf("no such page:%s\n", req.URL)
		http.Error(w, err, http.StatusNotFound)
	}
}

// kill $(lsof -t -i:5000)

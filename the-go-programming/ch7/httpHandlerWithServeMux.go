package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollar

type dollar float64

func main() {
	db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (d dollar) String() string {
	return fmt.Sprintf("$%2f", d)
}

func (db database) list(rw http.ResponseWriter, request *http.Request) {
	for item, price := range db {
		fmt.Fprintf(rw, "%s: %s\n", item, price)
	}
}

func (db database) price(rw http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	itemName := q.Get("item")
	price, ok := db[itemName]
	if !ok {
		err := fmt.Sprintf("not a valid item: %s", itemName)
		http.Error(rw, err, http.StatusNotFound)
		return
	}
	fmt.Fprintf(rw, "%2f", price)
}

func (db database) create(rw http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(rw, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var bodyJson struct {
		Price float32 `json:"price"`
		Name  string  `json:"name"`
	}
	err := json.NewDecoder(request.Body).Decode(&bodyJson)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	db[bodyJson.Name] = dollar(bodyJson.Price)
}

// curl http://localhost:8000/create --request POST --data '{"name":"book","price":3.99}'

func (db database) update(rw http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPut {
		http.Error(rw, "invalid request method", http.StatusMethodNotAllowed)
	}
	var body struct {
		Name  string
		Price float32
	}
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := db[body.Name]; !ok {
		http.Error(rw, "not a valid product", http.StatusBadRequest)
		return
	}

	db[body.Name] = dollar(body.Price)
	fmt.Fprintf(rw, "Updated product\n%s %s", body.Name, dollar(body.Price))
}

// curl http://localhost:8000/create --request POST --data '{"name":"book","price":3.99}'
// curl http://localhost:8000/update --request PUT --data '{"name":"book","price":2.99}'

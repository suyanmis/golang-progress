package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const url string = "https://api.github.com/issues"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	token := os.Getenv("GIT_TOKEN")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	query := req.URL.Query()
	query.Add("sort", "created")
	query.Add("since", time.Now().AddDate(-5, -1, 0).Format(time.RFC3339))
	req.URL.RawQuery = query.Encode()
	log.Println(req.URL.String())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println(string(decodeHttpResponseToJson(*resp)))

}

func decodeHttpResponseToJson(response http.Response) []byte {
	var result interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	prettyJSON, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return prettyJSON
}

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Transfer struct {
	Id          string
	Amount      int
	Created     int
	AvailableOn int `json:"available_on"`
	Currency    string
	Description string
	Fee         int
	Net         int
	Category    string
	Status      string
	Type        string
}

func main() {
	req, err := http.NewRequest("GET", "https://api.stripe.com/v1/balance_transactions", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("sk_test_51NmYi7E76ZgI9KtXBgSfN0qLCcvfMt8i8PbJYNqe4Fo5EB7V9grS91fyjwK4XhTpHDoqMz7Yovm7TTPurdat6t3U00KiUzT8EI", "")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got data: %v", string(body))

	var data struct {
		Object  string
		Data    []Transfer
		HasMore bool `json:"has_more"`
		Url     string
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	log.Printf("marshaled object: %+v", data)

	for _, transfer := range data.Data {
		log.Printf("result: %v", transfer)
	}

}

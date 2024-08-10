package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type JsonRes struct {
	Disclaimer string `json:"disclaimer"`
	License    string `json:"license"`
	Timestamp  int64  `json:"timestamp"`
	//Base is always USD for free api key
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
type ErrorRes struct {
	Error       bool   `json:"error"`
	Status      int64  `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func RequestMoreData() map[string]float64 {
	fmt.Println("Requesting more data...")
	res, err := RequestDataFromOpenExchange()
	if err != nil {
		log.Fatal(err)
	}

	mapExchange := res.Rates
	fmt.Println("Data updated")

	return mapExchange
}

func RequestDataFromOpenExchange() (*JsonRes, error) {
	API_KEY := os.Getenv("APP_ID")
	requestUrl := "https://openexchangerates.org/api/latest.json"
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
    header := "Token "+API_KEY
	request.Header.Add("Authorization", header)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errRes := ErrorRes{}
		json.NewDecoder(resp.Body).Decode(&errRes)
		fmt.Println(errRes)
		return nil, errors.New(errRes.Message)
	}

	jsonRes := JsonRes{}
	json.NewDecoder(resp.Body).Decode(&jsonRes)
	return &jsonRes, nil
}

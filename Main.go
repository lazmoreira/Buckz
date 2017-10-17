package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)

	if err != nil {
		panic(err)
	}
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

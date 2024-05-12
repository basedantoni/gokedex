package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Next != "" {
		url = c.Next
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	locationArea := LocationArea{}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Println(err)
	}

	c.Next = locationArea.Next
	c.Previous = url

	for _, v := range locationArea.Results {
		fmt.Println(v.Name)
	}

	return nil
}
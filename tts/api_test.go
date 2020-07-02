package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

//////////
// Please see my comments in README.md regarding tests. These tests are only
// examples of what would be a much larger test suite.
//////////

// TestNewShoeAndGetShoes tests both the newShoe and getShoes functions
// together.
func TestNewShoeAndGetShoes(t *testing.T) {
	var err error
	apiBasePath = "/api/v0"
	apiPort = 8080

	// Init new DB in memory.
	db, err = initDB(":memory:")
	if err != nil {
		t.Errorf("Error initializing test DB: %s", err)
		return
	}

	// Set up routes and start server
	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("%s/shoes", apiBasePath), newShoe).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/shoes", apiBasePath), getShoes).Methods("GET")
	go http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router)

	// POST a shoe.
	// 😂 https://stockx.com/nike-kyrie-5-spongebob
	spongebob := Shoe{
		Brand: "Nike",
		Name:  "Kyrie 5 Spongebob Squarepants",
	}
	spongebobBody, _ := json.Marshal(spongebob)
	spongebobResp, err := http.Post(
		fmt.Sprintf("http://localhost:%d%s/shoes", apiPort, apiBasePath),
		"application/json",
		bytes.NewBuffer(spongebobBody),
	)
	if err != nil {
		t.Errorf("Error POSTing spongebob: %s", err)
		return
	}
	defer spongebobResp.Body.Close()
	if spongebobResp.StatusCode != 201 {
		t.Errorf("Error POSTing shoe: Expected 201 response, got %d", spongebobResp.StatusCode)
		return
	}

	// POST another shoe.
	// I love this color. It would have matched my old car.
	// https://stockx.com/timberland-world-hiker-front-country-boot-supreme-orange
	orange := Shoe{
		Brand: "Timberland",
		Name:  "World Hiker Front Country Boot Supreme Orange",
	}
	orangeBody, _ := json.Marshal(orange)
	orangeResp, err := http.Post(
		fmt.Sprintf("http://localhost:%d%s/shoes", apiPort, apiBasePath),
		"application/json",
		bytes.NewBuffer(orangeBody),
	)
	if err != nil {
		t.Errorf("Error POSTing orange: %s", err)
		return
	}
	defer orangeResp.Body.Close()
	if orangeResp.StatusCode != 201 {
		t.Errorf("Error POSTing shoe: Expected 201 response, got %d", orangeResp.StatusCode)
		return
	}

	// GET two shoes.
	shoesResp, err := http.Get(fmt.Sprintf("http://localhost:%d%s/shoes", apiPort, apiBasePath))
	if err != nil {
		t.Errorf("Error GETing shoes: %s", err)
		return
	}
	defer shoesResp.Body.Close()
	if shoesResp.StatusCode != 200 {
		t.Errorf("Error GETing shoe: Expected 200 response, got %d", shoesResp.StatusCode)
		return
	}
	shoesRespBody, _ := ioutil.ReadAll(shoesResp.Body)
	var shoes []Shoe
	if err = json.Unmarshal(shoesRespBody, &shoes); err != nil {
		t.Errorf("Error unmarshaling JSON response body: %s", err)
		return
	}

	// make sure we have two shoes
	if len(shoes) != 2 {
		t.Errorf("Error reading shoes: Expected 2 shoes, got %d", len(shoes))
		return
	}
	// validate the Nike Spongebob
	if !((shoes[0].Brand == spongebob.Brand && shoes[0].Name == spongebob.Name) ||
		(shoes[1].Brand == spongebob.Brand && shoes[1].Name == spongebob.Name)) {
		t.Errorf("Error reading shoes: no spongebob")
	}
	// validate the Timberland Orange
	if !((shoes[0].Brand == orange.Brand && shoes[0].Name == orange.Name) ||
		(shoes[1].Brand == orange.Brand && shoes[1].Name == orange.Name)) {
		t.Errorf("Error reading shoes: no spongebob")
	}
}

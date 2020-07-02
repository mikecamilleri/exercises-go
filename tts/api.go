package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// newShoe is a handler function that creates a new shoe.
func newShoe(w http.ResponseWriter, r *http.Request) {
	// Should check headers here ("Content-Type" ...).

	// Decode the request body.
	var shoe Shoe
	json.NewDecoder(r.Body).Decode(&shoe)

	// Validate the shoe.
	if shoe.ID != 0 || shoe.Brand == "" || shoe.Name == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		// User error so DEBUG level. More info in message would be helpful.
		log.Print("DEBUG: bad request from user: ", shoe)
		return
	}

	// Write to the database.
	err := createShoe(shoe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		// Sending the raw error to the user here is a security risk, but for
		// this exercise I think it's fine and helpful.
		fmt.Fprintf(w, "%s", err)
		log.Print("ERROR: ", err) // should have been caught in validation
		return
	}

	// Write the response.
	w.WriteHeader(http.StatusCreated) // 201
}

// getShoes is a handler function that retrieves all of the shoes.
func getShoes(w http.ResponseWriter, r *http.Request) {
	// Should check headers here ("Accept" ...).

	// Read the shoes from the database.
	shoes, err := readShoes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		// Sending the raw error to the user here is a security risk, but for
		// this exercise I think it's fine and helpful.
		fmt.Fprintf(w, "%s", err)
		log.Print("ERROR: ", err)
		return
	}

	// Write the response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shoes)
}

// newTTSValueForShoe is a handler function that creates a new tts value for a
// shoe.
func newTTSValueForShoe(w http.ResponseWriter, r *http.Request) {
	// Should check headers here ("Content-Type" ...).

	// Decode the path.
	shoeID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// Maybe this should be 404 Not Found? I'm not sure if 404 is ever used
		// with methods other than GET.
		w.WriteHeader(http.StatusBadRequest) // 400
		// User error so DEBUG level. More info in message would be helpful.
		log.Print("DEBUG: bad request from user: id in url = ", mux.Vars(r)["id"])
		return
	}

	// Decode the request body.
	var tts TTS
	json.NewDecoder(r.Body).Decode(&tts)

	// Validate the tts object.
	// Because the shoe ID is included in the path, it should not be included
	// in the request body.
	if tts.ID != 0 || tts.ShoeID != 0 || tts.Value < 1 || tts.Value > 5 {
		w.WriteHeader(http.StatusBadRequest) // 400
		// User error so DEBUG level. More info in message would be helpful.
		log.Print("DEBUG: bad request from user: ", tts)
		return
	}

	// Add the shoe ID from the path to the tts object.
	tts.ShoeID = shoeID

	// Write to the databse.
	err = createTTSValue(tts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		// Sending the raw error to the user here is a security risk, but for
		// this exercise I think it's fine and helpful.
		fmt.Fprintf(w, "%s", err)
		log.Print("ERROR: ", err) // should have been caught in validation
		return
	}

	// Write the response.
	w.WriteHeader(http.StatusCreated) // 201
}

// getTTSAverageForShoe is a handler function that retrieves the average tts
// value for a shoe.
func getTTSAverageForShoe(w http.ResponseWriter, r *http.Request) {
	// Should check headers here ("Accept" ...).

	// Decode the path.
	shoeID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		// User error so DEBUG level. More info in message would be helpful.
		log.Print("DEBUG: bad request from user: id in url = ", mux.Vars(r)["id"])
		return
	}

	// Read and calculate the average tts value.
	ttsAvgValue, err := readTTSAverage(shoeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		// Sending the raw error to the user here is a security risk, but for
		// this exercise I think it's fine and helpful.
		fmt.Fprintf(w, "%s", err)
		log.Print("ERROR: ", err)
		return
	}

	// Create the object to be returned.
	averageTTS := AverageTTS{ShoeID: shoeID, Value: ttsAvgValue}

	// Write the response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(averageTTS)
}

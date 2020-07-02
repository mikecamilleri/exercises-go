package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	// SQLite driver for database/sql
	_ "github.com/mattn/go-sqlite3"
	// Aside from installing Postgres and creating the DB, switching to Postgres
	// would be as easy as using the driver below and changing the connection
	// string in sql.Open() in initDB().
	// _ "github.com/lib/pq"
)

var (
	configPath  = "./config.json"
	db          *sql.DB
	dbPath      string
	apiBasePath string
	apiPort     int
)

// Shoe represent a shoe (👟👟).
//
// Go source code is UTF-8 encoded which is fun becasue it lets us use emoji in
// our comments 🤣.
type Shoe struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

// TTS represents a single true-to-size (tts) value associated with a shoe.
//
// Values may be integers 1-5 inclusive.
type TTS struct {
	ID     int `json:"id"`
	ShoeID int `json:"shoeID"`
	Value  int `json:"value"`
}

// AverageTTS represents an average true-to-size (tts) value associated with a
// shoe.
//
// Values may be float32 1-5 inclusive.
type AverageTTS struct {
	ShoeID int     `json:"shoeID"`
	Value  float32 `json:"value"`
}

func main() {
	var err error

	// Read the configuration file and set global variables.
	err = setConfigFromFile(configPath)
	if err != nil {
		log.Fatal("FATAL: error reading configuration file: ", err)
	}

	// Connect to and set up database if not already set up (create tables etc.)
	log.Print("INFO: initializing database ...")
	db, err = initDB(dbPath)
	if err != nil {
		log.Fatal("FATAL: error initializing database: ", err)
	}
	defer db.Close()
	log.Print("INFO: database successfully initialized")

	// The routes below are very rudementary and just enough to accomplish what
	// is asked for in this assignment. There are some obvious enpoints missing
	// such as the ability to get a single shoe object by ID. I have designed
	// this, however, with extensibility in mind.
	router := mux.NewRouter()
	// create shoe
	router.HandleFunc(fmt.Sprintf("%s/shoes", apiBasePath), newShoe).Methods("POST")
	// get shoes
	router.HandleFunc(fmt.Sprintf("%s/shoes", apiBasePath), getShoes).Methods("GET")
	// create new tts value for shoe id
	router.HandleFunc(fmt.Sprintf("%s/shoes/{id}/tts", apiBasePath), newTTSValueForShoe).Methods("POST")
	// get tts average for shoe id
	router.HandleFunc(fmt.Sprintf("%s/shoes/{id}/tts", apiBasePath), getTTSAverageForShoe).Methods("GET")

	// Start the server and die if it fails.
	log.Print("INFO: serving http ...")
	err = http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router)
	if err != nil {
		log.Fatal("FATAL: error serving http: ", err)
	}

	os.Exit(1) // control should never reach this point
}

func setConfigFromFile(name string) error {
	var config map[string]string
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	json.NewDecoder(f).Decode(&config)

	if config["dbPath"] == "" {
		return errors.New("empty dbPath")
	}
	if config["apiBasePath"] == "" {
		return errors.New("empty apiBasePath")
	}
	if config["apiPort"] == "" {
		return errors.New("empty apiPort")
	}
	if _, err = strconv.Atoi(config["apiPort"]); err != nil {
		return errors.New("invalid apiPort")
	}

	dbPath = config["dbPath"]
	apiBasePath = config["apiBasePath"]
	apiPort, _ = strconv.Atoi(config["apiPort"])

	return nil
}

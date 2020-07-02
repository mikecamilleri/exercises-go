package main

import (
	"database/sql"

	// SQLite driver for database/sql
	_ "github.com/mattn/go-sqlite3"
	// Aside from installing Postgres and creating the DB, switching to Postgres
	// would be as easy as using the driver below and changing the connection
	// string in sql.Open() in initDB().
	// _ "github.com/lib/pq"
)

// initDB sets up the database connection and initializes the database if
// necessary.
//
// .Close() must be called on the returned *sql.DB object when it is no longer
// needed.
func initDB(path string) (*sql.DB, error) {
	var err error

	// "open" the database, which doesn't actually connect to it.
	db, err = sql.Open("sqlite3", path)

	// Connect to the DB by pinging it.
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	// Create required tables if they do not already exist.
	err = createTables()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// createTables creates the required database tables if they do not already
// exist.
func createTables() error {
	var err error

	sqlStatementShoes := `
		CREATE TABLE IF NOT EXISTS shoes (
			id INTEGER PRIMARY KEY, 
			brand TEXT NOT NULL, 
			name TEXT NOT NULL
		)`
	sqlStatementTTS := `
		CREATE TABLE IF NOT EXISTS tts (
			id INTEGER PRIMARY KEY, 
			shoe_id INT REFERENCES shoes(id) NOT NULL, 
			value INT CONSTRAINT one_through_five CHECK (value >= 1 AND value <= 5) NOT NULL
		)`

	_, err = db.Exec(sqlStatementShoes)
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlStatementTTS)
	if err != nil {
		return err
	}

	return nil
}

// createShoe allows a new shoe to be added to the database.
func createShoe(shoe Shoe) error {
	// db.Exec includes built in SQL injection protection when used this way!
	sqlStatement := `INSERT INTO shoes (brand, name) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, shoe.Brand, shoe.Name)
	if err != nil {
		return err
	}

	return nil
}

// readShoes returns all shoes in the database.
func readShoes() ([]Shoe, error) {
	// Query
	rows, err := db.Query("SELECT id, brand, name FROM shoes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Build the returned slice.
	var shoes []Shoe
	for rows.Next() {
		var shoe Shoe
		err = rows.Scan(&shoe.ID, &shoe.Brand, &shoe.Name)
		if err != nil {
			return nil, err
		}
		shoes = append(shoes, shoe)
	}

	// Get any additional errors.
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return shoes, nil
}

// createTTSValue allows for a TTS value to be added to the database.
func createTTSValue(tts TTS) error {
	// db.Exec includes built in SQL injection protection when used this way!
	sqlStatement := `INSERT INTO tts (shoe_id, value) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, tts.ShoeID, tts.Value)
	if err != nil {
		return err
	}
	return nil
}

// readTTSAverage queries for all TTS values associated with a shoe, calculates
// the mean, and returns a float32.
func readTTSAverage(shoeID int) (float32, error) {
	// Query TTS values for the shoe ID.
	// Query includes built in SQL injection protection when used this way!
	rows, err := db.Query("SELECT value FROM tts WHERE shoe_id=$1", shoeID)
	if err != nil {
		return 0.0, err
	}
	defer rows.Close()

	// Build the returned slice.
	var sum int
	var count int
	for rows.Next() {
		var n int
		err = rows.Scan(&n)
		if err != nil {
			return 0, err
		}
		sum += n
		count++
	}

	// Get any additional errors.
	err = rows.Err()
	if err != nil {
		return 0, err
	}

	// Check for no rows.
	if count == 0 {
		// Reusing this error from the sql package. There could probably be a
		// very long debate regarding whether this or creating my own error is
		// better form.
		return 0, sql.ErrNoRows
	}

	return (float32(sum) / float32(count)), nil
}

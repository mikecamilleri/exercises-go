package main

import "testing"

//////////
// Please see my comments in README.md regarding tests. These tests are only
// examples of what would be a much larger test suite.
//////////

// TestCreateShoeAndReadShoes tests both the createShoe and readShoes functions
// together.
func TestCreateShoeAndReadShoes(t *testing.T) {
	var err error
	var shoes []Shoe

	// Init new DB in memory.
	db, err = initDB(":memory:")
	if err != nil {
		t.Errorf("Error initializing test DB: %s", err)
		return
	}

	// Add a shoe.
	// 😂 https://stockx.com/nike-kyrie-5-spongebob
	spongebob := Shoe{
		Brand: "Nike",
		Name:  "Kyrie 5 Spongebob Squarepants",
	}
	if err = createShoe(spongebob); err != nil {
		t.Errorf("Error creating shoe: %s", err)
		return
	}

	// Add another shoe.
	// I love this color. It would have matched my old car.
	// https://stockx.com/timberland-world-hiker-front-country-boot-supreme-orange
	orange := Shoe{
		Brand: "Timberland",
		Name:  "World Hiker Front Country Boot Supreme Orange",
	}
	if err = createShoe(orange); err != nil {
		t.Errorf("Error creating shoe: %s", err)
		return
	}

	// Read two shoes.
	shoes, err = readShoes()
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

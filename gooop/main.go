package main

import "fmt"

type flyer interface {
	fly()
	crash()
}

func startFlight(f flyer) {
	fmt.Println("Starting flight ...")
	f.fly()
}

func causeCrash(f flyer) {
	fmt.Println("Causing crash ...")
	f.crash()
}

type meta struct {
	id   string
	name string
}

func (m *meta) printMeta() {
	fmt.Println("ID: " + m.id)
	fmt.Println("Name: " + m.name)
}

type alien struct {
	meta
	emoji rune
}

func (a *alien) fly() {
	fmt.Println("🛸")
}

func (a *alien) crash() {
	fmt.Println("💥")
}

func main() {
	fmt.Println("Hello Universe!")

	// composition and the "inheritance" of methods ...
	fmt.Println("I am an alien --")
	a := new(alien)
	a.id = "1"       // from the `meta` struct
	a.name = "Frank" // from the `meta` struct
	a.emoji = '👽'
	a.printMeta() // defined as a mehtod on the `meta` struct

	// now we us the `flyer` interface which `alien` implements ...
	fmt.Println("Let's Fly! --")
	startFlight(a)
	fmt.Println("Oh no! --")
	causeCrash(a)

	fmt.Println("Goodbye Universe.")
}

package main

import (
	"fmt"
)

type Animal interface {
	Says() string
	numberOfLeges() int
	//doesTail() bool // both dog and cat should has this fun implement
}

type Dog struct {
	name  string
	sound string
	legs  int
}

func (d *Dog) Says() string {
	return d.sound
}

func (d *Dog) numberOfLeges() int {
	return d.legs
}

type Cat struct {
	name    string
	sound   string
	legs    int
	hasTail bool
}

func (d *Cat) Says() string {
	return d.sound
}

func (d *Cat) numberOfLeges() int {
	return d.legs
}

func (c *Cat) doesTail() bool {
	return c.hasTail
}

func main() {
	dog := Dog{
		name:  "Dog",
		sound: "woof",
		legs:  4,
	}
	Riddle(&dog)

	cat := Cat{
		name:    "cat",
		sound:   "meow",
		legs:    4,
		hasTail: true,
	}
	// Riddle(cat)     // cant pass cat here even if having same info because we passed paramgeter type as Dog to function
	Riddle(&cat) // using interface we pass cat also to same function

}

func Riddle(a Animal) {
	fmt.Printf("animal sayas %s and has %d legs \n", a.Says(), a.numberOfLeges())
}

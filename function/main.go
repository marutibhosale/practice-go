package main

import "fmt"

type animals struct {
	name  string
	sound string
	legs  int
}

func (a *animals) says() { // passing struct to function  called methods
	fmt.Printf("%s says %s \n", a.name, a.sound)

}

func (a *animals) howManyLeags() {
	fmt.Printf("%s has %d legs \n", a.name, a.legs)

}

func main() {
	var dog animals
	dog.name = "dog"
	dog.sound = "boww"
	dog.legs = 4
	dog.says()
	dog.howManyLeags()

	var cat animals
	cat.name = "cat"
	cat.sound = "mayou"
	cat.legs = 4
	cat.says()
	cat.howManyLeags()

}

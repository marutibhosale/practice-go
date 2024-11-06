package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune // define channel

func main() {
	keyPressChan = make(chan rune)
	//keyPressChan <- 't'
	go listenForKeyPress() // go key word before funciton call its called goroutines

	fmt.Println("press any key, or q for quit")
	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}

}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You Pressed.", string(key))
	}

}

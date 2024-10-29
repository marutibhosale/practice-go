package main

import (
	// "bufio"
	// "fmt"
	// "os"
	// "strings"

	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

// go get -u github.com/eiannone/keyboard   used command to import third party package

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("-> ")

	// 	userInput, _ := reader.ReadString('\n')

	// 	userInput = strings.Replace(userInput, "\n", "", -1)

	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}
	// }

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key on keyboard and press ESC to quit")

	for {
		char, key, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal()
		}

		if key != 0 {
			fmt.Println("You Pressed ", char, key)
		} else {
			fmt.Println("You Pressed ", char)

		}

		if key == keyboard.KeyEsc {
			break
		}
	}
	fmt.Println("Program exsiting...")
}

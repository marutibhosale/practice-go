package yesno

import (
	"log"
	"fmt"
	"github.com/eiannone/keyboard"
)

func GetYesOrNo (s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'n' || char == 'N' {
			return false
		}else if char == 'y' || char == 'Y' {
			return true
		} else {
			fmt.Println("press y,Y,n,N only")
		}

	}
}
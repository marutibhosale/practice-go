package main

// go its not oop lang its simillar like compositions
import (
	"bufio"
	"fmt"
	yesno "myapp/yes_no"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	name             string
	age              int
	flovoriateNumber float64
	wantDogs         bool
}

func main() {
	userdata := User{ // like this also we pass valune to user defined variables like struct
		name:             "maruti",
		age:              3,
		flovoriateNumber: 32.09,
		wantDogs:         true,
	}
	println(userdata)
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.name = readString("what is your name")
	user.age = readInt("what is your age")
	user.flovoriateNumber = readFloat("enter your flovorite number")
	user.wantDogs = yesno.GetYesOrNo("Do you want dogs?")

	//fmt.Println("your name is", user.name,". and you are",user.age, "year old.")
	//fmt.Println(fmt.Sprintf("your name is %s. and you are %d years old.",user.name,user.age))
	fmt.Printf("your name is %s, you are %d years old and your favorite number is %.2f, wantDogs: %t \n", user.name, user.age, user.flovoriateNumber, user.wantDogs)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("enter correct name")
		} else {
			return userInput
		}
	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("enter correct age")
		} else {
			return num
		}
	}

}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("enter float number")
		} else {
			return num
		}
	}

}

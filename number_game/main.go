package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press enter when ready"

func main () {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number")
	fmt.Println("======================")
	fmt.Println("")

	fmt.Println("Think of number between 1 to 10",prompt)
	reader.ReadString('\n')

	fmt.Println("multiply number by", firstNumber , prompt)
	reader.ReadString('\n')

	fmt.Println("multiply nuumber by" , secondNumber, prompt, )
    reader.ReadString('\n')

	fmt.Println("divide number by number you thought", prompt)
	reader.ReadString('\n')

	fmt.Println("subract the number", subtraction, prompt)
	reader.ReadString('\n')

	// answer
    answer = firstNumber * secondNumber - subtraction
	fmt.Println("Amswer is", answer)

}
package main

import (
	"fmt"
	packagetwo "myapp/packageone" //
)

func main() {
	fmt.Println(packagetwo.PublicVar)
	packagetwo.Exported()
}

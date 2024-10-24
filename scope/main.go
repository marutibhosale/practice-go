package main

import (
	"fmt"
	"myapp/packageone" //
)

var myVar = "main-package-var"

func main() {
	fmt.Println(packageone.PublicVar)
	packageone.Exported()


	blockVar := "main-block-var"

	packageone.PrintMe(myVar,blockVar)
}

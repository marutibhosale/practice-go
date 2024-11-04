package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string // slice
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "horse")
	animals = append(animals, "cow")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)

	}

	fmt.Println(sort.StringsAreSorted(animals))

	//fmt.Printf("Type of x: %T\n", animals)

	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//delete(intMap, "three")
	fmt.Println(intMap)

	ele, ok := intMap["three"]
	if ok {
		fmt.Println(ele,"exsists")
	} else {
		fmt.Println(ele ,"doesnt exists")
	}

}

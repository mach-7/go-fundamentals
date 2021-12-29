package main

import "fmt"

func main() {

	myMap := map[string]string{

		"name":  "Lionel Messi",
		"club":  "Barcelona",
		"shirt": "10",
	}

	// maps can expand dynamically as opposed to static structs
	myMap["position"] = "Forward"

	fmt.Printf("%+v\n", myMap)

	printMap(myMap)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}

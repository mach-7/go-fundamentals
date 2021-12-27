package main

import "fmt"

type location struct {
	latitude  float64
	longitude float64
}

type person struct {
	firstname string
	lastname  string
	gender    string
	dob       string
	location  location
}

func main() {
	//mylocation := location{32.005, 71.044}
	person := person{
		firstname: "Albus",
		lastname:  "Dumbledore",
		gender:    "Male",
		dob:       "Unknown",
		location: location{
			latitude:  34.5,
			longitude: 84.0,
		},
	}
	fmt.Printf("%+v", person)
}

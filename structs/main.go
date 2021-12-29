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
	p := person{
		firstname: "Albus",
		lastname:  "Dumbledore",
		gender:    "Male",
		dob:       "Unknown",
		location: location{
			latitude:  34.5,
			longitude: 84.0,
		},
	}
	p.print()
	// pass by value does not change the value of the original struct
	p.updateName("Cornelius")
	p.print()

	//Pass by reference on the other hand results in modification of the original struct
	pointerP := &p
	pointerP.updateNamePointer("Cornelius")
	p.print()
}

func (p person) updateName(newName string) {
	p.firstname = newName
}

func (p *person) updateNamePointer(newName string) {
	(*p).firstname = newName
}

func (p person) print() {
	fmt.Printf("\n%+v", p)
}

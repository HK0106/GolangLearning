package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{firstName: "Jim", lastName: "Doe"}

	jim.updatePersonName("Jimmy")
	jim.print()
}

func (p *person) updatePersonName(name string) {
	(*p).firstName = name
}
func (p person) print() {
	fmt.Printf("%v", p)
}

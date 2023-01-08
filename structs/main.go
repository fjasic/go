package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	filip := person{firstName: "Flip", lastName: "Jasic", contactInfo: contactInfo{email: "filipjasic8@gmail.com", zip: 21400}}
	filip.updateName("Filip")
	filip.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

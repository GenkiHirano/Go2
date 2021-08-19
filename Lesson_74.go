package main

import "fmt"

type Person struct {
	firstName string
}

func (a Person) name() string {
	return a.firstName
}

type User struct {
	Person
}

func (a User) name() string {
	return a.firstName
}

func main() {
	bob := Person{"Bob"}
	mike := User{}
	mike.firstName = "Mike"

	fmt.Println(bob.name())
	fmt.Println(mike.name())
}

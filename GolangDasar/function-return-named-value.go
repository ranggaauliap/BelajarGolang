package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Rangga"
	middleName = "Aulia"
	lastName = "Putra AAA"

	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}

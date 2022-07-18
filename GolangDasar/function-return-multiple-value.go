package main

import "fmt"

func getFullName() (string, string) {
	return "Rangga", "Aulia Putra"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	//==> menghiraikan return value

	NamaPertama, _ := getFullName()
	fmt.Println(NamaPertama)

}

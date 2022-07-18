package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, ", My Name is", customer.Name)
}

func main() {
	var rangga Customer
	rangga.Name = "Rangga Aulia Putra"
	rangga.Address = "Indonesia"
	rangga.Age = 25

	rangga.sayHai("Rangga AP")

	/*fmt.Println(rangga.Name)
	fmt.Println(rangga.Address)
	fmt.Println(rangga.Age)

	// struct literals

	aulia := Customer{
		Name:    "Aulia",
		Address: "Indonesia",
		Age:     25,
	}
	fmt.Println(aulia)

	putra := Customer{"Putra", "Indonesia", 25}
	fmt.Println(putra)
	*/
}

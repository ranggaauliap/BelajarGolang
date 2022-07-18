package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChanceCountrToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Kemayoran", "Jakarta Pusat", "Indonesia"}
	//var address4 *Address = &Address{"Kemayoran", "Jakarta Pusat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Surakarta"
	*address2 = Address{"Jogja", "DIY", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Surakarta",
		Province: "Jawa Tengah",
		Country:  "",
	}

	var alamatPointer *Address = &alamat
	ChanceCountrToIndonesia(alamatPointer)
	fmt.Println(alamat)
}

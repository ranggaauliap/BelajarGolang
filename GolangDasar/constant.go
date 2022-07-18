package main

import "fmt"

func main() {
	const firstName string = "Rangga"
	const lastname = "Aulia Putra"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastname)
	fmt.Println(value)

	const (
		namadepan    string = "Rangga"
		namabelakang        = "Aulia Putra"
		nilai               = 1000
	)

	fmt.Println(namadepan)
	fmt.Println(namabelakang)
	fmt.Println(nilai)
}

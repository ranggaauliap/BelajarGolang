package main

import "fmt"

func main() {
	name := "Rangga"

	switch name {
	case "Rangga":
		fmt.Println("Hello Rangga")
		fmt.Println("Hello Rangga")
	case "Aulia":
		fmt.Println("Hello Aulia")
		fmt.Println("Hello Aulia")
	case "Putra":
		fmt.Println("Hello Putra")
		fmt.Println("Hello Putra")
	default:
		fmt.Println("Hello ", name)
		fmt.Println("Hello ", name)
	}

	switch length := len(name); length > 6 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Oke")
	}

	length := len(name)
	switch {
	case length > 6:
		fmt.Println("Nama Terlalu Panjang")
	case length < 1:
		fmt.Println("Nama Terlalu Pendek")
	default:
		fmt.Println("Oke")
	}
}

package main

import "fmt"

func main() {
	var name = "Rangga"
	var length = len(name)

	if name == "Rangga" {
		fmt.Println("Hello Rangga")
	} else if name == "Aulia" {
		fmt.Println("Hello Aulia")
	} else {
		fmt.Println("Hello ", name)
	}

	if length > 6 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Oke")
	}
}

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	slice := []string{"Rangga", "Aulia", "Putra"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//==> For Statement

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// ==> For Range

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Rangga"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

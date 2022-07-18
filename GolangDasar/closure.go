package main

import "fmt"

func main() {
	name := "Rangga"
	counter := 0

	increment := func() {
		name := "Aulia"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}

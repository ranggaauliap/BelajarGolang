package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Rangga"
	names[1] = "Aulia"
	names[2] = "Putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		96,
		97,
		98,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}

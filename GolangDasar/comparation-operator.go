package main

import "fmt"

func main() {
	var name1 = "Rangga"
	var name2 = "Rangga"

	var result bool = name1 == name2
	fmt.Println(result)

	var nama1 = "Rangga"
	var nama2 = "Adi"

	var result1 bool = nama1 > nama2
	fmt.Println(result1)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}

package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPrangga NoKTP = "131345135132525"
	var marriedStatus Married = false
	fmt.Println(noKTPrangga)
	fmt.Println(marriedStatus)
}

package main

import "fmt"

<<<<<<< HEAD
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	result := goodbye("Rangga")
	fmt.Println(result)
	fmt.Println(goodbye("Rangga"))
=======
func getodBye(name string) string {
	return "od Bye " + name
}

func main() {
	odbye := getodBye
	result := odbye("Rangga")
	fmt.Println(result)
	fmt.Println(odbye("Rangga"))
>>>>>>> 6bdfcc015eae824d20ef758197cad274be9f0a0f
}

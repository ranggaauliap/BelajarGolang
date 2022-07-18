package main

import "fmt"

func random() interface{} {
<<<<<<< HEAD
	return "Rangga"
=======
	return 100
>>>>>>> 6bdfcc015eae824d20ef758197cad274be9f0a0f
}

func main() {
	var result interface{} = random()
<<<<<<< HEAD
	var resultString string = result.(string)
	fmt.Println(resultString)
=======
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
>>>>>>> 6bdfcc015eae824d20ef758197cad274be9f0a0f
}

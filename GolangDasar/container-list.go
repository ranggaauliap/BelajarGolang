package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Rangga")
	data.PushBack("Aulia")
	data.PushBack("Putra")
	data.PushFront("Raden")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	//dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}

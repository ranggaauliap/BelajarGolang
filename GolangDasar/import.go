package main

import (
	"BelajarGolang/GolangDasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Rangga")
	// helper.sayGoodbye("Rangga") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //error
}

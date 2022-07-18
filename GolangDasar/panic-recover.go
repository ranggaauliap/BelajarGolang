package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error message dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}

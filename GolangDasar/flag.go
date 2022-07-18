package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put Your Database Host")
	var user *string = flag.String("user", "root", "Put Your Database User")
	var password *string = flag.String("password", "root", "Put Your Database Password")
	var number *int = flag.Int("number", 100, "Put Your Number")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("User :", *user)
	fmt.Println("Password :", *password)
	fmt.Println("Number :", *number)

}

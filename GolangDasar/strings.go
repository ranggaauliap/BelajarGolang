package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rangga Aulia Putra", "Rangga"))
	fmt.Println(strings.Split("Rangga Aulia Putra", "a"))
	fmt.Println(strings.ToLower("Rangga Aulia Putra"))
	fmt.Println(strings.ToUpper("Rangga Aulia Putra"))
	fmt.Println(strings.Trim("   Rangga Aulia Putra    ", " "))
	fmt.Println(strings.ReplaceAll("Rangga Aulia Putra", "Putra", "Santoso"))
}

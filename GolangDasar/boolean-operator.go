package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 90
	var lulusNilaiAbsensi bool = absensi >= 80

	var lulus = lulusNilaiAbsensi && lulusNilaiAkhir
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 90 && absensi >= 80)
}

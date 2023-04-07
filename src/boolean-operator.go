package main

import "fmt"

func main(){
	var nilaiUjian = 85
	var nilaiAbsensi = 75

	var lulusUjian = nilaiUjian >= 80
	var lulusAbsensi = nilaiAbsensi >= 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
	
	fmt.Println(nilaiUjian >= 80 && nilaiAbsensi >= 80)
	fmt.Println(nilaiUjian >= 80 || nilaiAbsensi >= 80)

}
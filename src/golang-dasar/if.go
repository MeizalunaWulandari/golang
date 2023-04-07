package main

import "fmt"

func main(){
	var name = "Luna"

	if name == "Luna" {
		fmt.Println("Halo Luna")
	} else if name == "Meiza" {
		fmt.Println("Halo Meiza")
	} else {
		fmt.Println("Halo", name, "boleh kenalan ga?")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	}else{
		fmt.Println("Nama sudah bener")
	}

	// Short statement 
	if length := len(name) ; length > 5 {
		fmt.Println("Nama terlalu panjang")
	}else{
		fmt.Println("Nama sudah bener")
	}

}
/*
	Pada else if, else akan dieksekusi jika semua kondisi tidak memenuhi
*/
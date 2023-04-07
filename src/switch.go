package main

import "fmt"

func main(){
	name := "Meizaluna Wulandari"

	switch name {
	case "Meiza":
		fmt.Println("Halo Meiza")
	case "Luna":
		fmt.Println("Halo Luna")
	case "Andini":
		fmt.Println("Halo Meiza")
	default:
		fmt.Println("Hai, boleh kenalan ga?")
	}

	// Short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah bener")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length < 3:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah bener")
	}
	
}

/*
	default blok yang akan di eksekusi jika semua syarat tidak terpenuhi
	pada default tidak perlu memasukkan expression lagi
	Go-Lang memungkinkan untuk membuat switch tanpa statement
*/
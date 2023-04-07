package main

import "fmt"

func main(){
	var name string
	name = "Meizaluna"
	fmt.Println(name)

	name = "Meizaluna Wulandari"
	fmt.Println(name)

	var motherName = "Wulan"
	fmt.Println(motherName)

	motherName = "Andini Puspa Wulandari"
	fmt.Println(motherName)

	var age int8 = 38
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Meizaluna"
		lastName = "Wulandari"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}

/*
	pada golang variable yang telah dibuat harus dipakai, jika tidak dipakai akan terjadi error
	jika langsung mengisi variable maka tidak perlu mendefinisikan type
	`var` tidak wajib saat menginisiasi variable keyword tersebut dapat diganti dengan :=
	pada golang dapat mendeklarasikan multiple variable dengan keyword var()
*/
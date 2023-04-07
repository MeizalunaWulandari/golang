package main

import "fmt"

type Biodata struct{
	Name, Address string
	Age 		  int
	Married 	  bool
}

func main(){
	var luna Biodata
	luna.Name = "Meizaluna Wulandari"
	luna.Address = "Banjarmasin"
	luna.Age = 16
	luna.Married = false

	fmt.Println(luna)
	fmt.Println(luna.Name)
	fmt.Println(luna.Address)
	fmt.Println(luna.Age)
	fmt.Println(luna.Married)

	// Struct Literal

	andini := Biodata{
		Name : "Andini Puspa Wulandari",
		Address : "Banjarmasin",
		Age : 36,
		Married : true,
	}

	fmt.Println(andini)
	fmt.Println(andini.Name)
	fmt.Println(andini.Address)
	fmt.Println(andini.Age)
	fmt.Println(andini.Married)

	rizka := Biodata{
		Name : "Rizka Aulia",
		Address : "Banjarmasin",
		Age : 16,
		Married : false,
	}
	fmt.Println(rizka)
	fmt.Println(rizka.Name)
	fmt.Println(rizka.Address)
	fmt.Println(rizka.Age)
	fmt.Println(rizka.Married)

}

/**
 * Struct adalah sebuah template yang digunakan untuk menggabungkan nol atau lebih type data yang digunakan dalam saru kesatuan
 * Struct biasa representasi dari data dalam program kita
 * Struct disimpan dalam field
 * Sederhananya struct adalah kumpulan dari field
 * struct bisa dibuat dengan uppercase seperti NamaStruct, bukan seperti namaStruct, namastruct Namastruct
 * Struct tidak bisa digunakan langsung
 * Namun kita bisa membuat data/object dari struct yang kita buat
*/
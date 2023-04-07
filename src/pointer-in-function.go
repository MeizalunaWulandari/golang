package main

import "fmt"

type Address struct{
	City, Provice, Country string
}

func changeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
	
	var alamat = Address {
		City : "Banjarmasin",
		Provice: "Kalimantan Selatan",
		Country: "",
	}
	// changeCountryToIndonesia(alamat)
	// fmt.Println(alamat)
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}

/**
 * Pada parameter funct tambahkan * dan pada saat pemanggilan tambahkan &
 * Ketika membuat stuct yang besar disarankan menggunakan pointer agar menghemat resource
*/
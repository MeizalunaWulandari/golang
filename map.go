package main

import "fmt"

func main(){
	person := map[string]string{
		"name": "Meizaluna Wulandari",
		"address": "Banjarmasin",
	}
	person["education"] = "SMK/SMK Frater don Bosco"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["education"])
	fmt.Println(len(person))

	book := make(map[string]string)
	book["judul"] = "Belajar Go-Lang"
	book["pembeli"] = "Meizaluna Wulandari"
	book["penjual"] = "Andini"
	book["ups"] = "Salah"

	delete(book, "ups")
	fmt.Println(book)
	
}

/*
	singkatnya map sepeerti asosiatif array pada pemerogram PHP, yaitu kumpulan dan value
	Tidak seperti array atau slice yang memiliki jumlah maksimal length atau capacity yang harus di definisikan sebelum membuatnya
	map dapat dibuat sebanyak-banyaknya dengan syarat keynya unique, namun jika membuat key yang sama otomatis value dari key yang lama akan dai ganti dengan yang baru
	perlu diperhatikan untuk membuat map pasangan key dan value di pisahkan dengan : (titik dua)
	function pada map
		len(map) untuk mendapat jumlah data di map
		map[key] mengambil data di map dengan key
		map [key] = value mengubah data di map dengan key 
		make(map[typeKey]typrValue) membuat map baru
		delete(map, key) menghapus data di map dengan value

*/
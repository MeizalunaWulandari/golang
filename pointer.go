package main

import "fmt"

type Address struct{
	City, Provice, Country string
}

func main(){
	address1 := Address{"Banjarmasin", "Kalimantan Selatan", "Indonesia"}

	// Tanpa Pointer
	address2 := address1
	address2.City = "Balikpapan" //Tidak berubah

	// fmt.Println(address1)
	// fmt.Println(address2)

	// Dengan Pointer 
	address3 := &address1
	address3.City = "Balikpapan" //Tidak berubah
	// fmt.Println(address1)
	// fmt.Println(address3)

	*address3 = Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	fmt.Println(address4)


}

/**
 * Saat mengganti variable di golang sebenarnya yang dilakukan oleh golang iyalah menyalin value
 * var := value1
 * var2 := var
 * Yang sebenarnya terjadi iyalah
 * var1 := value1
 * var2 := value1
 * sehingga jika merubah data di variable 2 data di variable 1 tidak akan terpenggaruh
 * Pointer sendiri adalah kemampuan untuk membuat reference ke lokasi data yang sama tanpa menduplikasinya
 * var1 := value1
 * var2 := var1
 * dengan pointer data tesebut tidak akan duplikasi seperti contoh atas melainkan merefence ke variable 1
 * var1 := value1
 * var2 := var1
 * Untuk membuat variable dengan nilain pointer kita bisa menggunakan operator & diikuti dengan variable lain
 * var1 := value1
 * var2 := &var1
 * 
 * Selain & dan * untuk membuat pointer juga bisa dengan function new
 * namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
 * 
*/
package main

import "fmt"

func main(){
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(len(months))

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(months)
	fmt.Println(slice1)

	slice1[0] = "Slice Diubah"
	fmt.Println(months)
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Luna")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Meizaluna"
	newSlice[1] = "Wulandari"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

}

/*
	Slice adalah potongan array
	slice mirip seperti array hanya saja ukurannya bisa diubah
	function pada slice
		len() = untuk mendapat panjang
		cap() = untuk mendapat kapasitas
		append() = (slice, data) = untuk membuat slice baru dengan menambah data ke posisi terakhir, jika kapastitas sudah penuh maka akan mebuat array baru
		make([]typeData, length, capacity) = membuat slice baru
		copy(destination, source) = menyalin slice dari source ke destination, usahakan capacity dan length sama jika tidak sama akan terpotong
	hasil fmt.Println(cap(slice1)) = 8 karna dimulai dari mei, dihitung sampai akhir (desember)
	saat merubah array slice juga ikut berubah
	begitu juga saat slice diubah array juga ikut berubah
	slice1[0] = index 0 merujuk pada awal array di potong yaitu 4 (mei)

	Hati-hati saat membuat array karna bisa jadi yang dibuat itu slice,
		[...]string  // ini array
		[]string 	 // ini slice

*/
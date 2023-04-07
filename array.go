package main

import "fmt"

func main(){
	var names [2]string
		names[0] = "Meizaluna"
		names[1] = "Wulandari"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(len(names))

	var values = [3]int{
		20,
		30,
		30,
	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))

	var lagi = [10]int{}
	fmt.Println(len(lagi))
}

/*
	Sebelum membuat array harus menentukan jumlah maksimal data yang akan disimpan
	Untuk membuat array secara langsung gunakan `{}` dan diakhiri dengan koma `,`
	func len() untuk mengetahui panjang array
*/
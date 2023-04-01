package main

import "fmt"

func main(){
	const firstName = "Meizaluna"
	const lastName = "Wulandari"
	const umur = 16

	const(
		motherName = "Wulan"
		age = 38
	)

	// error
	/*
		firstName = "Andini"
		lastName = "Wulandari"
	*/
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(umur)
	fmt.Println(motherName)
	fmt.Println(age)
}

/*
	Constant adalah variable yang nilainya tidak bisa diubah setelah pertama kali diberi nilai
	untuk membuat constant sama dengan membuat variable namun dengan keyword `const`
	saat membuat constant wajib untuk langsung mengisi datanya
	berbeda dengan variable constant yang telah dibuat tidak akan error jika tidak digunakan 
	untuk membuat multiple constant sama seperti variable namun dengan keyword `const`
*/